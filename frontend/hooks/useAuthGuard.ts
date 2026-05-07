import { useEffect } from "react";
import { useRouter } from "next/navigation";

export const useAuthGuard = () => {
  const router = useRouter();

  useEffect(() => {
    const userId = localStorage.getItem("piggy_user_id");
    if (!userId) {
      router.push("/login");
    }
  }, [router]);
};
