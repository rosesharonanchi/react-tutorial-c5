"use client";

import Navbar from "@/components/navbar";
import TransactionsList from "@/components/TransactionsList";
import { useSearchParams } from "next/navigation";
import { useGetTransactions } from "@/hooks/useFetchTransactions";

export default function TransactionsClient() {
  const searchParams = useSearchParams();
  const typeParam = searchParams.get("type");
  const type =
    typeParam === "saving" || typeParam === "withdrawal"
      ? typeParam
      : undefined;

  const { transactions, loading } = useGetTransactions({
    type,
  });

  return (
    <div className="flex flex-col min-h-screen">
      <Navbar />

      <main className="max-w-3xl mx-auto p-6">
        <h1>All Transactions</h1>

        {loading ? (
          <p>Loading...</p>
        ) : (
          <TransactionsList transactions={transactions || []} />
        )}
      </main>
    </div>
  );
}
