// "use client";
// import TransactionsList from "@/components/TransactionsList";
// import Navbar from "@/components/navbar";
// import { useGetTransactions } from "@/hooks/useFetchTransactions";
// import { GetTransactionsParamsType, TransactionType } from "@/types/interfaces";
// import { useSearchParams } from "next/navigation";
// import React from "react";

// function TransactionsPage() {
// 	const query = useSearchParams();
// 	const type = query.get("type");
// 	const {transactions, loading} = useGetTransactions({
// 		type: type,
// 	} as GetTransactionsParamsType);

// 	console.log("IN Transactions: ", transactions);

// 	return (
//     <div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
//       <Navbar />
//       <main className="w-full max-w-3xl mx-auto px-6 py-10">
//         <h1 className="text-xl font-semibold text-slate-700 dark:text-slate-300 mb-4">
//           All Transactions
//         </h1>
//         {loading ? (
//           <p>Loading...</p>
//         ) : (
//           <TransactionsList transactions={transactions || []} />
//         )}
//       </main>
//     </div>
//   );
// }

// export default TransactionsPage;
import { Suspense } from "react";
import TransactionsClient from "./TransactionsClient";

export default function Page() {
  return (
    <Suspense fallback={<p>Loading...</p>}>
      <TransactionsClient />
    </Suspense>
  );
}
