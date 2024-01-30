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

export async function GetFacilities(park: string) {
  var rawData = await getData(links[park]);
  if (rawData == null) {
    return [{ FacilityName: "failed", FacilityID: "0" } as Facility];
  }

  return JSON.parse(rawData) as Facility[];
}

async function getData(url: string) {
  const res = await fetch(url, {
    next: { revalidate: 60 },
    headers: {
      "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) ",
    },
  });

  if (!res.ok) {
    console.error("Failed to fetch data: ", res.statusText);
    console.error("Failed to fetch data: ", url);
    return null;
  }

  return res.text();
}

interface Dictionary<T> {
  [Key: string]: T;
}

const links: Dictionary<string> = {
  tdl: "https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json",
  tds: "https://www.tokyodisneyresort.jp/_/realtime/tds_attraction.json",
};
