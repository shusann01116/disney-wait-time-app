export interface Link {
  attraction: string;
  greeting: string;
}

const links = {
  tdl_attraction:
    "https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json",
  tds_attraction:
    "https://www.tokyodisneyresort.jp/_/realtime/tds_attraction.json",
  tdl_greeting: "https://www.tokyodisneyresort.jp/_/realtime/tdl_greeting.json",
  tds_greeting: "https://www.tokyodisneyresort.jp/_/realtime/tds_greeting.json",
} as const;

export function getLink(park: "tdl" | "tds"): Link | null {
  switch (park) {
    case "tdl":
      return {
        attraction: links.tdl_attraction,
        greeting: links.tdl_greeting,
      };
    case "tds":
      return {
        attraction: links.tds_attraction,
        greeting: links.tds_greeting,
      };
    default:
      return null;
  }
}

export interface FacilityOperatingHours {
  OperatingHoursFromDate: string;
  OperatingHoursFrom: string;
  OperatingHoursToDate: string;
  OperatingHoursTo: string;
  SunsetFlg: boolean;
  OperatingStatusCD: string | null;
  OperatingStatus: string | null;
  OperatingChgFlg: boolean;
}

export interface Facility {
  FacilityID: string;
  FacilityName: string;
  FacilityKanaName: string;
  NewFlg: boolean;
  FacilityURLSP: string | null;
  FacilityStatusCD: string | null;
  FacilityStatus: string | null;
  StandbyTime: boolean | string | null;
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

  // greeting
  greeting: FacilityGreeting;
  cameraman?: FacilityCameraman;
}

export interface OperatingHours {
  OperatingHoursFromDate: string;
  OperatingHoursFrom: string;
  OperatingHoursToDate: string;
  OperatingHoursTo: string;
  OperatingStatusCD: string | null;
  OperatingStatus: string | null;
  OperatingChgFlg: boolean;
  SunsetFlg: boolean;
}

export interface FacilityGreeting {
  FacilityID: string;
  FacilityName: string;
  FacilityKanaName: string;
  NewFlg: boolean;
  AreaJName: string;
  AreaMName: string;
  FacilityURLSP: string | null;
  FacilityStatusCD: string | null;
  FacilityStatus: string | null;
  StandbyTime: string | null;
  operatinghours: OperatingHours[];
  UseStandbyTimeStyle: boolean;
  UpdateTime: string;
}

export interface FacilityCameraman {
  FacilityID: string;
  FacilityName: string;
  FacilityKanaName: string;
  FacilityStatusCD: string;
  FacilityStatus: string;
  operatinghours: OperatingHours[] | null;
  UpdateTime: string;
}

export interface Area {
  AreaJName: string;
  AreaMName: string;
  Facility: Facility[];
}

export interface Greetings {
  [key: string]: Area;
}
