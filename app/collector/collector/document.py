import json
import logging

import requests
from attraction import Attraction

logger = logging.getLogger(__name__)


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
