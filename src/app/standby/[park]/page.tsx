import { ParkType } from "@/lib/fetcher";
import { StandbyList } from "./standby-list";
import { Suspense } from "react";

export const dynamicParams = false;
export const revalidate = 60;

export default async function Page({
  params,
}: {
  params: Promise<{ park: string }>;
}) {
  const park = (await params).park;
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <StandbyList park={park as ParkType} />
    </Suspense>
  );
}

export async function generateStaticParams() {
  return [{ park: "tdl" }, { park: "tds" }];
}
