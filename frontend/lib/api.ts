import axios from "axios";
import { BASE_URL } from "./constants";

export const getUserData = async () => {
  const storedId = localStorage.getItem("piggy_user_id");
  if (!storedId || storedId === "undefined") return null;

  try {
    const url = `${BASE_URL}/api/v1/users/${storedId}`;
    const res = await axios.get(url);
    return res.data; // This will contain total_savings and total_withdrawals
  } catch (error) {
    console.error("Error fetching user stats:", error);
    return null;
  }
};
