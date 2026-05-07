import { TransactionType } from "@/types/interfaces";
import axios from "axios";
import { BASE_URL } from "@/lib/constants";


export interface ResponseType {
  success: boolean;
  error: any;
}


export const saveTransaction = async (
  payload: TransactionType,
): Promise<ResponseType> => {
  console.log("Fetch function executed!");

  const userId = localStorage.getItem("piggy_user_id");
  try {
    const res = await axios.post(BASE_URL + "/api/v1/transactions", payload, {
      headers: {
        "X-User-ID": userId || "",
      },
    });
    return {
      success: true,
      error: null,
    };
  } catch (err) {
    return {
      success: false,
      error: err,
    };
  }
};
