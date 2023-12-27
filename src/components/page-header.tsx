"use client";

import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";
import Link from "next/link";

export default function PageHeader() {
  return (
    <header className="sticky right-0 top-0 flex justify-center border-b border-border/40 bg-background/95 backdrop-blur-lg">
      <div className="container flex max-w-screen-2xl items-center">
        <Link href="/" className="mr-6 font-sans font-bold">
          <span>Disney App</span>
        </Link>
        <NavigationMenu>
          <NavigationMenuList>
            <NavigationMenuItem>
              <Link href="/standby" legacyBehavior passHref>
                <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                  Standby
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>
      </div>
    </header>
  );
}
