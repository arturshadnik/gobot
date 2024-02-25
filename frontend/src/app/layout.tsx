import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from '@/components/shared/Header'
import { Suspense } from "react";
import { AuthProvider } from "@/lib/auth/authContext";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Goot",
  description: "Simple AI chatbot built with a Go backend.",
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
          </Suspense>{children}
        </body>
      </AuthProvider>
    </html>
  );
}
