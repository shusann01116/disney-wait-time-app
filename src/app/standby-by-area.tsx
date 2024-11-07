"use client";

import { StandbyCard } from "@/components/ui/standby-card";
import { getArea } from "@/lib/attractions";
import { FacilityResp } from "@/lib/fetcher";
import { useState } from "react";

type Props = {
  standbyList: FacilityResp[];
};

export default function StandbyByArea({ standbyList }: Props) {
  const [favoriteList, setFavoriteList] = useState<Record<string, boolean>>(
    standbyList.reduce(
      (acc, facility) => {
        acc[facility.id] = false;
        return acc;
      },
      {} as Record<string, boolean>,
    ),
  );

  const groupedByArea = standbyList.reduce(
    (acc, facility) => {
      acc[getArea(facility.id)] = acc[getArea(facility.id)] || [];
      acc[getArea(facility.id)].push(facility);
      return acc;
    },
    {} as Record<string, FacilityResp[]>,
  );

  return (
    <>
      {Object.entries(groupedByArea).map(([areaName, facilities]) => (
        <div key={areaName} className="space-y-4">
          <h2 className="text-2xl font-bold">{areaName}</h2>
          <ul className="space-y-4">
            {facilities.map((facility) => (
              <li key={facility.id}>
                <StandbyCard
                  facility={facility}
                  size="sm"
                  onFavorite={(id) => {
                    setFavoriteList((prev) => ({ ...prev, [id]: !prev[id] }));
                  }}
                  isFavorite={favoriteList[facility.id] ?? false}
                />
              </li>
            ))}
          </ul>
        </div>
      ))}
    </>
  );
}
