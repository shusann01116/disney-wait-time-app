import { Facility } from "./types";
import { FacilityResp } from "./fetcher";

export function toFacilityRespFromAttraction(f: Facility): FacilityResp {
  return {
    id: f.FacilityID,
    name: f.FacilityName,
    operatingStatus: {
      id: f.OperatingStatusCD ?? "",
      name: f.OperatingStatus ?? "",
    },
    operatingHour: {
      from: new Date(`${f.OperatingHoursFromDate} ${f.OperatingHoursFrom}:00`),
      to: new Date(`${f.OperatingHoursToDate} ${f.OperatingHoursTo}:00`),
    },
    standbyTime:
      typeof f.StandbyTime === "string" ? parseInt(f.StandbyTime, 10) : 0,
    updatedAt: new Date(f.UpdateTime),
  };
}

export function toFacilityRespFromGreeting(
  f: Facility,
): FacilityResp | undefined {
  const greeting = f.greeting;
  if (!greeting) return undefined;

  const operatingHours = greeting.operatinghours?.[0] || {};

  return {
    id: greeting.FacilityID,
    name: greeting.FacilityName,
    operatingStatus: {
      id: operatingHours.OperatingStatusCD ?? "",
      name: operatingHours.OperatingStatus ?? "",
    },
    operatingHour: {
      from: new Date(
        `${operatingHours.OperatingHoursFromDate} ${operatingHours.OperatingHoursFrom}:00`,
      ),
      to: new Date(
        `${operatingHours.OperatingHoursToDate} ${operatingHours.OperatingHoursTo}:00`,
      ),
    },
    standbyTime:
      typeof greeting.StandbyTime === "string"
        ? parseInt(greeting.StandbyTime, 10)
        : 0,
    updatedAt: new Date(greeting.UpdateTime),
  };
}
