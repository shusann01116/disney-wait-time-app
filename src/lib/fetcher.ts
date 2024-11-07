"use server";

import { Facility, getLink, Greetings } from "@/lib/types";
import {
  toFacilityRespFromAttraction,
  toFacilityRespFromGreeting,
} from "./facility";

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
  return facilities.map(toFacilityRespFromAttraction);
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

  const parsedData = JSON.parse(greetings) as Greetings;
  const facilities: FacilityResp[] = [];

  Object.values(parsedData).forEach((area) => {
    area.Facility.forEach((facility) => {
      const facilityResp = toFacilityRespFromGreeting(facility);
      if (facilityResp) {
        facilities.push(facilityResp);
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
