import { cn } from "@/lib/utils";
import Link from "next/link";
import { Poiret_One } from "next/font/google";

const poiret = Poiret_One({
  subsets: ["latin"],
  weight: ["400"],
  variable: "--font-poiret",
});

export default function PageHeader() {
  return (
    <header className="flex justify-center bg-background/50 px-4 py-1 backdrop-blur-lg">
      <span className="w-full max-w-screen-lg">
        <Link href="/" className={cn(poiret.className, "text-4xl font-bold")}>
          Disney Portal
        </Link>
      </span>
    </header>
  );
}
