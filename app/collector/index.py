import json
import logging
import os
from datetime import datetime, timedelta, timezone

import boto3
import collector.document as document
from collector.attraction import Attraction, AttractionRepository

logger = logging.getLogger(__name__)

DYNAMODB_TABLENAME = os.getenv("DYNAMODB_TABLENAME")
if DYNAMODB_TABLENAME is None:
    logger.error("DYNAMODB_TABLENAME is not set.")
    raise Exception()

LOGGING_LEVEL = os.getenv("LOGGING_LEVEL")
if LOGGING_LEVEL is None:
    LOGGING_LEVEL = "INFO"
logger.setLevel(LOGGING_LEVEL)


base_url = "https://www.tokyodisneyresort.jp"

paths = [
    "/_/realtime/tdl_attraction.json",
    "/_/realtime/tds_attraction.json",
    "/_/realtime/tdl_greeting.json",
    "/_/realtime/tds_greeting.json",
]


def handler(event, context):
    logger.info(event)
    main()
    return {"statusCode": 200, "body": json.dumps({"message": "ok"})}


def main() -> None:
    repo = AttractionRepository(boto3.client("dynamodb"), DYNAMODB_TABLENAME)
    attractions: list[Attraction] | None = None
    for path in paths:
        attractions = document.parse_document(
            document.fetch_document(
                base_url
                + path
                + "?"
                + str(int(datetime.now(timezone(timedelta(hours=9))).timestamp())),
            ),
        )

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
