import { getStandbys } from "@/lib/fetcher";
import StandbyByArea from "./standby-by-area";

export default async function Page() {
  const [standbyListTdl, standbyListTds] = await Promise.all([
    getStandbys("tdl"),
    getStandbys("tds"),
  ]);
  const standbyList = [...standbyListTdl, ...standbyListTds];
  return <StandbyByArea standbyList={standbyList} />;
}
