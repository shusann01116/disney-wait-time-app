import { badgeVariants } from "@/components/ui/badge";
import { GetFacilities } from "@/lib/fetcher";
import { cn } from "@/lib/utils";
import Link from "next/link";

export default async function Home({ params }: { params: { park: string } }) {
  const facilities = await GetFacilities(params.park);

  return (
    <>
      <div className="flex gap-4 py-4">
        <Link
          href="/standby/tdl"
          className={cn(
            badgeVariants({ variant: "outline" }),
            "bg-muted text-muted-foreground",
            {
              "bg-inherit text-foreground": params.park === "tdl",
            },
          )}
        >
          DisneyLand
        </Link>
        <Link
          href="/standby/tds"
          className={cn(
            badgeVariants({ variant: "outline" }),
            "bg-muted text-muted-foreground",
            {
              "bg-inherit text-foreground": params.park === "tds",
            },
          )}
        >
          DisneySea
        </Link>
      </div>
      <ul className="w-full max-w-3xl divide-y">
        {facilities.map((facility) => {
          return (
            <li
              key={facility.FacilityID}
              className="overflow-hidden bg-background py-4"
            >
              <div className="text-wrap text-xl font-bold text-foreground">
                {facility.FacilityName}
              </div>
              <div className="mt-2 flex">
                <div className="items-center text-sm text-secondary-foreground">
                  {facility.OperatingStatus}
                </div>
                <div className="flex-grow text-right text-sm text-secondary-foreground">
                  {facility.StandbyTime === null ? (
                    <></>
                  ) : (
                    <span className={cn("text-sm text-secondary-foreground")}>
                      {facility.StandbyTime} 分待ち
                    </span>
                  )}
                </div>
              </div>
            </li>
          );
        })}
      </ul>
      <div className="mt-4"></div>
    </>
  );
}
