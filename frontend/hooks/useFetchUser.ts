import { useState, useEffect } from "react";
import { getUserData } from "@/lib/api";

export function useFetchUserStats() {
  const [stats, setStats] = useState({
    total_savings: "0",
    total_withdrawals: "0",
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getStats = async () => {
      const data = await getUserData();
      if (data) {
        setStats({
          total_savings: data.total_savings,
          total_withdrawals: data.total_withdrawals,
        });
      }
      setLoading(false);
    };
    getStats();
  }, []);

  return { stats, loading };
}
