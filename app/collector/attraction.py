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
