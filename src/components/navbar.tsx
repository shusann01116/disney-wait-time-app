"use client";

import {
  Link,
  Navbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
} from "@nextui-org/react";
import { usePathname } from "next/navigation";

export default function Top() {
  const pathname = usePathname();
  return (
    <Navbar position="sticky">
      <NavbarBrand>
        <Link href="/" className="ml-2 text-2xl font-bold text-black">
          Disney App
        </Link>
      </NavbarBrand>
      <NavbarContent className="flex gap-4 justify-center ">
        <NavbarItem isActive={pathname === "/"}>
          <Link href="/" className="text-black" aria-current="page">
            ホーム
          </Link>
        </NavbarItem>
        <NavbarItem isActive={pathname.includes("/standby")}>
          <Link href="/standby/tdl" className="text-black">
            待ち時間
          </Link>
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}
