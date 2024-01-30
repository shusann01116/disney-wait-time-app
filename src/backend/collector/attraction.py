from datetime import datetime, timedelta, timezone
from decimal import Decimal
from typing import Any, Literal, Mapping, Sequence, Set, Union

from mypy_boto3_dynamodb import DynamoDBClient
from mypy_boto3_dynamodb.type_defs import AttributeValueTypeDef


class Attraction:
    def __init__(
        self,
        FacilityID,
        FacilityName,
        FacilityKanaName,
        NewFlg,
        FacilityURLSP,
        FacilityStatusCD,
        FacilityStatus,
        StandbyTime,
        OperatingHoursFromDate,
        OperatingHoursFrom,
        OperatingHoursToDate,
        OperatingHoursTo,
        OperatingStatusCD,
        OperatingStatus,
        SunsetFlg,
        Fsflg,
        FsStatusflg,
        FsStatus,
        FsStatusCD,
        FsStatusStartDate,
        FsStatusStartTime,
        FsStatusEndDate,
        FsStatusEndTime,
        UseLimitFlg,
        UseStandbyTimeStyle,
        OperatingChgFlg,
        UpdateTime,
        operatingHours,
    ):
        self.FacilityID = FacilityID
        self.FacilityName = FacilityName
        self.FacilityKanaName = FacilityKanaName
        self.NewFlg = NewFlg
        self.FacilityURLSP = FacilityURLSP
        self.FacilityStatusCD = FacilityStatusCD
        self.FacilityStatus = FacilityStatus
        self.StandbyTime = StandbyTime
        self.OperatingHoursFromDate = OperatingHoursFromDate
        self.OperatingHoursFrom = OperatingHoursFrom
        self.OperatingHoursToDate = OperatingHoursToDate
        self.OperatingHoursTo = OperatingHoursTo
        self.OperatingStatusCD = OperatingStatusCD
        self.OperatingStatus = OperatingStatus
        self.SunsetFlg = SunsetFlg
        self.Fsflg = Fsflg
        self.FsStatusflg = FsStatusflg
        self.FsStatus = FsStatus
        self.FsStatusCD = FsStatusCD
        self.FsStatusStartDate = FsStatusStartDate
        self.FsStatusStartTime = FsStatusStartTime
        self.FsStatusEndDate = FsStatusEndDate
        self.FsStatusEndTime = FsStatusEndTime
        self.UseLimitFlg = UseLimitFlg
        self.UseStandbyTimeStyle = UseStandbyTimeStyle
        self.OperatingChgFlg = OperatingChgFlg
        self.UpdateTime = UpdateTime
        self.operatingHours = operatingHours

    def __str__(self) -> str:
        return f"{self.FacilityName} - {self.StandbyTime} min ({self.OperatingStatus})"


class AttractionRepository:
    def __init__(self, client: DynamoDBClient, table_name: str):
        self.client: DynamoDBClient = client
        self.table_name = table_name

    def put_item(self, attraction: Attraction) -> None:
        try:
            item = self._to_dynamo_item(attraction)
            self.client.put_item(TableName=self.table_name, Item=item)
        except Exception as e:
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
        dt = self._parse_to_datetime(
            updateTime=attraction.UpdateTime,
            operatingHoursFromDate=attraction.OperatingHoursFromDate,
        )
        return {
            "attraction_id": {"S": attraction.FacilityID},  # PK
            "timestamp": {"S": dt.isoformat()},  # SK
            "date": {"S": dt.strftime("%Y-%m-%d")},  # GPK-1
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

    def _parse_to_datetime(
        self,
        *,
        updateTime: str,
        operatingHoursFromDate: str,
        timeDelta=9,
    ) -> datetime:
        # 23:21
        (hour, minute) = updateTime.split(":")
        (hour, minute) = (int(hour), int(minute))

        # 20230823
        (year, month, day) = (
            int(operatingHoursFromDate[0:4]),
            int(operatingHoursFromDate[4:6]),
            int(operatingHoursFromDate[6:8]),
        )

        if not 0 <= hour <= 24:
            raise Exception("hour is not in the range of 0-24")
        if not 0 <= minute <= 60:
            raise Exception("minute is not in the range of 0-60")
        if not 1 <= month <= 12:
            raise Exception("month is not in the range of 1-12")
        if not 1 <= day <= 31:
            raise Exception("day is not in the range of 1-31")

        return datetime(
            year=year,
            month=month,
            day=day,
            hour=hour,
            minute=minute,
            tzinfo=timezone(timedelta(hours=timeDelta)),
        )
