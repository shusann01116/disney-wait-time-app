import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { cn } from "@/lib/utils";
import Header from "@/components/page-header";
import { ThemeProvider } from "@/components/theme-provider";

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
        <ThemeProvider>
          <div className="container flex flex-col mx-auto items-center justify-center">
            {children}
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
