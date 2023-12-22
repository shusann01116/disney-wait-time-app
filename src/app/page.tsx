"use clinet";

import { GetFacilities } from "@/lib/fetcher";
import { cn } from "@/lib/utils";
import { Button, Navbar, NavbarBrand } from "@nextui-org/react";
import { revalidatePath } from "next/cache";

export default async function Home() {
  const facilities = await GetFacilities();

  return (
    <>
      {/* Navbar */}
      <Navbar position="sticky">
        <NavbarBrand>
          <p className="text-2xl font-bold">待ち時間一覧</p>
        </NavbarBrand>
      </Navbar>
      {/* Cards */}
      <div className="mt-4 px-4 w-full max-w-[1024px] flex flex-col mx-auto items-center">
        <ul className="columns-3xs lg:columns-3 gap-3">
          {facilities.map((facility) => {
            return (
              <li
                key={facility.FacilityID}
                className="bg-gray-50 drop-shadow-lg border border-slate-700 w-full min-w-unit-sm py-2 px-4 rounded-md my-2 first:mt-0 last:md-0 overflow-auto"
              >
                <div>
                  <div className="text-wrap text-xl font-bold">
                    {facility.FacilityName}
                  </div>
                  <div className="flex mt-2">
                    <div className="text-sm text-slate-500  items-center">
                      {facility.OperatingStatus}
                    </div>
                    <div className="text-sm text-slate-500 text-right flex-grow">
                      {facility.StandbyTime === null ? (
                        <></>
                      ) : (
                        <span className={cn("text-slate-700 text-sm")}>
                          {facility.StandbyTime} 分待ち
                        </span>
                      )}
                    </div>
                  </div>
                </div>
              </li>
            );
          })}
        </ul>
      </div>
      <div className="mt-4"></div>
    </>
  );
}
