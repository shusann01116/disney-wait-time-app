import json
import logging
import os
from collections.abc import Mapping
from datetime import datetime, timedelta, timezone
from decimal import Decimal
from typing import Any, Literal, Sequence, Set, Union

import boto3
import requests
from attraction import Attraction
from mypy_boto3_dynamodb.client import DynamoDBClient
from mypy_boto3_dynamodb.type_defs import AttributeValueTypeDef

logger = logging.getLogger(__name__)

DYNAMODB_TABLENAME = os.getenv("DYNAMODB_TABLENAME")
if DYNAMODB_TABLENAME is None:
    logger.error("DYNAMODB_TABLENAME is not set.")
    raise Exception()


base_url = "https://www.tokyodisneyresort.jp"

paths = [
    "/_/realtime/tdl_attraction.json",
    "/_/realtime/tds_attraction.json",
]


def handler(event, context):
    logger.info(event)
    main()
    return {"statusCode": 200, "body": json.dumps({"message": "ok"})}


class AttractionRepository:
    def __init__(self, client: DynamoDBClient, table_name: str):
        self.client: DynamoDBClient = client
        self.table_name = table_name

    def put_item(self, attraction: Attraction) -> None:
        try:
            item = self._to_dynamo_item(attraction)
            self.client.put_item(TableName=self.table_name, Item=item)
        except Exception as e:
            logger.error(e)
            raise e

    def _to_dynamo_item(
        self,
        attraction: Attraction,
    ) -> Mapping[
        str,
        Union[
            AttributeValueTypeDef,
            Union[
                bytes,
                bytearray,
                str,
                int,
                Decimal,
                bool,
                Set[int],
                Set[Decimal],
                Set[str],
                Set[bytes],
                Set[bytearray],
                Sequence[Any],
                Mapping[str, Any],
                None,
            ],
        ],
    ]:
        return {
            "attraction_id": {"S": attraction.FacilityID},  # PK
            "timestamp": {"S": self._print_iso(attraction.UpdateTime)},  # SK
            "name": {"S": self._get_str(attraction.FacilityName)},
            "status_id": {"S": self._get_str(attraction.OperatingStatusCD)},
            "status": {"S": self._get_str(attraction.OperatingStatus)},
            "wait_time": {"S": self._get_str(attraction.StandbyTime)},
        }

    def _get_str(self, value: str | Literal[False] | None) -> str:
        if value is None:
            return ""
        if value is False:
            return ""
        return value

    def _print_iso(self, str_time: str) -> str:
        (hour, minute) = str_time.split(":")
        return (
            datetime.now(timezone(timedelta(hours=9)))
            .replace(hour=int(hour), minute=int(minute), second=0, microsecond=0)
            .replace(tzinfo=timezone.utc)
            .isoformat(sep="T")
        )


def fetch_document(url: str, *, params: dict | None = None) -> str:
    headers = {"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "}
    return requests.get(url, headers=headers, params=params).text


def parse_document(document: str) -> list[Attraction]:
    result = []

    for a in json.loads(document):
        result.append(Attraction(**a))

    return result


def is_park_open(base_url: str) -> bool:
    path = "/view_interface.php"
    params = {
        "blockId": 94199,
        "pageBlockId": 5805799,
        "cmd": "openStatus",
        "rt": 1,
    }
    url = base_url + path
    return json.loads(fetch_document(url, params=params))["close_flg"] == "1"


def main() -> None:
    if not is_park_open(base_url):
        logger.info("Park is closed")
        return

    time_stamp = str(int(datetime.now(timezone(timedelta(hours=9))).timestamp()))

    attractions: list[Attraction] | None = None
    for path in paths:
        attractions = parse_document(fetch_document(base_url + path + "?" + time_stamp))

    if attractions is None:
        logger.warn("No information is fetched")
        return

    repo = AttractionRepository(boto3.client("dynamodb"), DYNAMODB_TABLENAME)
    for att in attractions:
        # TODO: make this loop async when the performance came up as an issue

        # Skip if any of the required fields are None for DynamoDB
        if att.FacilityID is None:  # PK
            logger.warn("FacilityID is None")
            continue
        if att.OperatingStatusCD is None:  # SK
            logger.warn("OperatingStatusCD is None")
            continue

        repo.put_item(att)


if __name__ == "__main__":
    main()
