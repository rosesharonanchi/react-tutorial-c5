import { GetTransactionsParamsType, TransactionType } from "@/types/interfaces";
import axios from "axios";
// export interface GetTransactionsRes {
//   transactions: TransactionType[];
//   error: any;
// }
import { BASE_URL } from "@/lib/constants";
export const getAllTransactions = async (
	query: GetTransactionsParamsType,
): Promise<TransactionType[]> => {

	const queries = [];
	if (query.size) {
		queries.push("size=" + query.size);
	}

	if (query.type) {
		queries.push("type=" + query.type);
	}

	const url =
    BASE_URL +
    "/transactions" +
    (queries.length > 0 ? "?" : "") +
    queries.join("&");

	try {
		const storedId = localStorage.getItem("piggy_user_id");
    const userId = storedId === "undefined" || !storedId ? "" : storedId;
		const res = await axios.get(url, {
			headers: {
				"X-User-ID": userId,
			},
		});
		

		return  res.data || [];
  }
   
    
	catch (error) {
    console.error("API Error:", error);
		return [];
	}
};



