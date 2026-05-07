import { useState } from "react";
import { useRouter } from "next/navigation";
import { registerRequest, loginRequest } from "@/api/auth";

export const useAuth = () => {
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  const signup = async (formData: any) => {
    setIsLoading(true);
    try {
      const res = await registerRequest(formData);
      // const data = await res.json();
      const userId = res.user_id || res.id || res.ID;

      if (userId) {
        localStorage.setItem("piggy_user_id", userId.toString());
        router.push("/");
      } else {
        alert("Account created, but no User ID received. Try logging in.");
      }
    } catch (err) {
      alert("Connection error");
    } finally {
      setIsLoading(false);
    }
  };

  // LOGIN
  const login = async (formData: any) => {
    setIsLoading(true);
    try {
      const res = await loginRequest(formData);

      // const data = await res.json();

      // if (res.ok) {
      const userId = res.user_id || res.id || res.ID;

      if (userId) {
        localStorage.setItem("piggy_user_id", userId.toString());
        router.push("/");
      } else {
        alert("Login successful, but no User ID received");
      }
      // } else {
      //   alert(res.error || "Login failed");
      // }
    } catch (err :any) {
      const errorMsg = err.response?.data?.error || "Invalid email or password";
      alert(errorMsg);

      alert("Server response was not valid JSON. Check Go terminal.");
    } finally {
      setIsLoading(false);
    }
  };

  return { signup, login, isLoading };
};
