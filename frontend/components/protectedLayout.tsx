"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function ProtectedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const [isAuthorized, setIsAuthorized] = useState(false);

  useEffect(() => {
    const userId = localStorage.getItem("piggy_user_id");

    if (!userId) {
      router.push("/login");
    } else {
      setIsAuthorized(true);
    }
  }, [router]);

  // Prevent UI flicker while checking localStorage
  if (!isAuthorized) {
    return null;
  }

  return <>{children}</>;
}
