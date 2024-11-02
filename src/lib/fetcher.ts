"use server";

import { Facility, getLink } from "@/lib/types";

export type FacilityResp = {
  id: string;
  name: string;
  operatingStatus: { id: string; name: string };
  operatingHour: { from: Date; to: Date };
  standbyTime: number;
  updatedAt: Date;
};

export type ParkType = "tdl" | "tds";

export async function getStandbys(park: ParkType): Promise<FacilityResp[]> {
  const attractions = await getAttractions(park);
  if (attractions === null) {
    return [];
  }

  const greetings = await getGreetings(park);
  if (greetings === null) {
    return [];
  }

  return [...attractions, ...greetings];
}

export async function getAttractions(park: ParkType): Promise<FacilityResp[]> {
  const link = getLink(park);
  if (link == null) {
    return [];
  }

  const attractions = await getData(link?.attraction);
  if (attractions == null) {
    return [];
  }

  const facilities = JSON.parse(attractions) as Facility[];

  return facilities.map((f) => ({
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
  }));
}

export async function getGreetings(park: ParkType): Promise<FacilityResp[]> {
  const link = getLink(park);
  if (link == null) {
    return [];
  }

  const greetings = await getData(link.greeting);
  if (greetings == null) {
    return [];
  }

  const parsedData = JSON.parse(greetings);
  const facilities: FacilityResp[] = [];

  Object.values(parsedData).forEach((area) => {
    // @ts-expect-error we know the type of Facility is FacilityWrapper
    area.Facility.forEach((facilityWrapper) => {
      const greeting = facilityWrapper.greeting;
      if (greeting) {
        const operatingHours = greeting.operatinghours?.[0] || {};
        facilities.push({
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
        });
      }
    });
  });

  return facilities;
}

async function getData(url: string) {
  const res = await fetch(url, {
    next: { revalidate: 60 },
    headers: {
      "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) ",
    },
  });

  if (!res.ok) {
    return null;
  }

  return res.text();
}
