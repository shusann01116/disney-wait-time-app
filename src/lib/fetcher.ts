"use server";

type FacilityOperatingHours = {
  OperatingHoursFromDate: string;
  OperatingHoursFrom: string;
  OperatingHoursToDate: string;
  OperatingHoursTo: string;
  SunsetFlg: boolean;
  OperatingStatusCD: string | null;
  OperatingStatus: string | null;
  OperatingChgFlg: boolean;
};

type Facility = {
  FacilityID: string;
  FacilityName: string;
  FacilityKanaName: string;
  NewFlg: boolean;
  FacilityURLSP: string | null;
  FacilityStatusCD: string | null;
  FacilityStatus: string | null;
  StandbyTime: string | null;
  OperatingHoursFromDate: string;
  OperatingHoursFrom: string;
  OperatingHoursToDate: string;
  OperatingHoursTo: string;
  OperatingStatusCD: string;
  OperatingStatus: string;
  SunsetFlg: boolean;
  Fsflg: boolean;
  FsStatusflg: string | null;
  FsStatus: string | null;
  FsStatusCD: string | null;
  FsStatusStartDate: string | null;
  FsStatusStartTime: string | null;
  FsStatusEndDate: string | null;
  FsStatusEndTime: string | null;
  UseLimitFlg: boolean;
  UseStandbyTimeStyle: boolean;
  OperatingChgFlg: boolean;
  UpdateTime: string;
  operatingHours: FacilityOperatingHours[];
};

export async function GetFacilities() {
  const rawData = await fetch(link, { next: { revalidate: 60 } }).then((res) =>
    res.text(),
  );
  const facilities: Facility[] = JSON.parse(rawData);
  return facilities;
}

const link = "https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json";
