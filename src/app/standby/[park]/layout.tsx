import { Button, Link } from "@nextui-org/react";
import { TbBuildingCastle, TbVolcano } from "react-icons/tb";
export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="md:flex">
      <div className="flex gap-4 md:flex-col md:mr-2 mb-4">
        <Button
          as={Link}
          href="/standby/tdl"
          className="bg-orange-200 border border-slate-700 shadow-lg"
        >
          Tokyo Disneyland
        </Button>
        <Button
          as={Link}
          href="/standby/tds"
          className="bg-sky-200 border border-slate-700 shadow-lg"
        >
          Tokyo DisneySea
        </Button>
      </div>
      {children}
    </div>
  );
}
