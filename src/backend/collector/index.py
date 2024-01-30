import json
import logging
import os
from datetime import datetime, timedelta, timezone

import boto3
import requests
from attraction import Attraction, AttractionRepository

logger = logging.getLogger(__name__)

DYNAMODB_TABLENAME = os.getenv("DYNAMODB_TABLENAME")
if DYNAMODB_TABLENAME is None:
    logger.error("DYNAMODB_TABLENAME is not set.")
    raise Exception()

S3_BUCKETNAME = os.getenv("S3_BUCKETNAME")
if S3_BUCKETNAME is None:
    logger.error("S3_BUCKETNAME is not set.")
    raise Exception()

LOGGING_LEVEL = os.getenv("LOGGING_LEVEL")
if LOGGING_LEVEL is None:
    LOGGING_LEVEL = "INFO"
logger.setLevel(LOGGING_LEVEL)


base_url = "https://www.tokyodisneyresort.jp"

paths = {
    "tdl_attraction": "/_/realtime/tdl_attraction.json",
    "tds_attraction": "/_/realtime/tds_attraction.json",
    "tdl_greeting": "/_/realtime/tdl_greeting.json",
    "tds_greeting": "/_/realtime/tds_greeting.json",
    "tdl_parade_show": "/_/realtime/tdl_parade_show.json",
    "tds_parade_show": "/_/realtime/tds_parade_show.json",
}


def handler(event, context):
    logger.info(event)
    main()
    return {"statusCode": 200, "body": json.dumps({"message": "ok"})}


def fetch_document(url: str, *, params: dict | None = None) -> str:
    headers = {"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "}
    docs = requests.get(url, headers=headers, params=params).text
    logger.info(f"fetched {len(docs)} bytes")
    return docs


def parse_document(document: str) -> list[Attraction]:
    result = []

    for a in json.loads(document):
        result.append(Attraction(**a))

    logger.info(f"parsed {len(result)} attractions")
    return result


def main() -> None:
    repo = AttractionRepository(boto3.client("dynamodb"), DYNAMODB_TABLENAME)
    s3 = boto3.client("s3")
    now = datetime.now(timezone(timedelta(hours=9)))
    attractions: list[Attraction] | None = None
    for key, path in paths.items():
        logger.info(f"fetching {path}")
        document = fetch_document(
            base_url + path + "?" + str(int(now.timestamp())),
        )
        s3.put_object(
            Body=document.encode("utf-8"),
            Bucket=S3_BUCKETNAME,
            Key="/".join(
                [
                    now.strftime("%Y%m%d"),
                    f'{now.strftime("%H%M")}_{key}.json',
                ]
            ),
        )

        attractions = parse_document(document)

        if attractions is None:
            logger.warn("no information is fetched")
            return

        for att in attractions:
            # TODO: make this loop async when the performance came up as an issue
            logger.debug(f"att: {att}")

            # Skip if any of the required fields are None for DynamoDB
            if att.FacilityID is None:  # PK
                logger.warn("FacilityID is None")
                logger.warn(f"skipping: {att}")
                continue
            if att.OperatingStatusCD is None:  # SK
                logger.warn("OperatingStatusCD is None")
                logger.warn(f"skipping: {att}")
                continue

            repo.put_item(att)

        logger.info(f"finished to put {len(attractions)} items")

    logger.info("done")


if __name__ == "__main__":
    main()
