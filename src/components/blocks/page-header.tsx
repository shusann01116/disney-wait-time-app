import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";
import { cn } from "@/lib/utils";
import Link from "next/link";

const menuItems = [
  {
    label: "Standby",
    href: "/standby",
  },
];

export default function PageHeader() {
  return (
    <header className="sticky left-0 top-0 mx-8 flex justify-center bg-background/50 py-1 backdrop-blur-lg">
      <span className="flex w-full max-w-screen-lg items-baseline gap-x-4">
        <Link href="/standby" className="font-sans font-bold">
          Disney App
        </Link>
        <NavigationMenu>
          <NavigationMenuList>
            {menuItems.map((item) => (
              <NavigationMenuItem key={item.href}>
                <NavigationMenuLink
                  className={cn(
                    navigationMenuTriggerStyle(),
                    "bg-transparent hover:underline hover:underline-offset-1",
                  )}
                  href={item.href}
                >
                  {item.label}
                </NavigationMenuLink>
              </NavigationMenuItem>
            ))}
          </NavigationMenuList>
        </NavigationMenu>
      </span>
    </header>
  );
}
