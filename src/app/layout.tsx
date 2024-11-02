import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { cn } from "@/lib/utils";
import PageHeader from "@/components/blocks/page-header";
import { ThemeProvider } from "@/components/providers/theme-provider";

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
    <html lang="en">
      <body className={cn(inter.className)}>
        <ThemeProvider attribute="class" defaultTheme="light" enableSystem>
          <PageHeader />
          <div className="container mx-auto flex flex-col items-center justify-center">
            {children}
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
