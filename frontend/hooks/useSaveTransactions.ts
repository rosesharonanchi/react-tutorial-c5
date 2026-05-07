import { useState } from "react";
import { saveTransaction } from "@/api/create-transaction";
import { TransactionType } from "@/types/interfaces";

export const useSaveTransaction = () => {
  const [loading, setLoading] = useState(false);

  const save = async (data: TransactionType) => {
    setLoading(true);
    const result = await saveTransaction(data);
    setLoading(false);
    return result;
  };

  return { save, loading };
};
