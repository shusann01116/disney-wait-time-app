import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Providers } from "./providers";
import { cn } from "@/lib/utils";
import Top from "@/components/navbar";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Disney App",
  description: "Dashboard app makes your disney life easier.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="black">
      <body className={cn(inter.className, "items-center justify-center")}>
        <Providers>
          <>
            <Top />
            <div className="mt-4 px-4 w-full max-w-[1024px] flex flex-col mx-auto items-center justify-center">
              {children}
            </div>
          </>
        </Providers>
      </body>
    </html>
  );
}
