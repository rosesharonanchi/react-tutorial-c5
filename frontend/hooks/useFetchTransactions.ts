import { useEffect, useState } from "react";
import { getAllTransactions } from "@/api/get-transactions";
import { GetTransactionsParamsType, TransactionType } from "@/types/interfaces";

export const useGetTransactions = (params: GetTransactionsParamsType) => {
  const [transactions, setTransactions] = useState<TransactionType[]>([]);
  const [loading, setLoading] = useState(true);
  const [totalSavings, setTotalSavings] = useState("0")
  const [totalWithdrawal, setTotalWithdrawal] = useState("0")
  useEffect(() => {
    const fetch = async () => {
      setLoading(true);
      try {
        const res = await getAllTransactions(params);
        setTransactions(res);
      } catch (error) {
        setTransactions([]);
      } finally {
        setLoading(false);
      }
    };
    fetch();
  }, [params.size, params.type]);

  return {
    transactions, 
    loading,
  };
};
