import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from '@/components/shared/Header'
import { Suspense } from "react";
import { AuthProvider } from "@/lib/auth/authContext";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Goot",
  description: "Gen AI capture the flag",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <AuthProvider>
        <body className={inter.className}>
          <Suspense fallback="...">
            <Header />
          {children}</Suspense>
        </body>
      </AuthProvider>
    </html>
  );
}
