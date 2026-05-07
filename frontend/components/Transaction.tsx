import { TransactionType } from "@/types/interfaces";
import React from "react";
export type { TransactionType as TransactionProps };

function Transaction({ amount, type, reason, createdAt }: TransactionType) {
	const isSaving = type === "saving";
	return (
		<tr className="border-b border-slate-100 dark:border-slate-700 hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors">
			<td className="py-3 px-4 font-medium text-slate-800 dark:text-slate-200">
				{amount} CFA
			</td>
			<td className="py-3 px-4">
				<span
					className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
						isSaving
							? "bg-green-100 text-green-700 dark:bg-green-900/40 dark:text-green-400"
							: "bg-red-100 text-red-700 dark:bg-red-900/40 dark:text-red-400"
					}`}
				>
					{isSaving ? "↑ Saving" : "↓ Withdrawal"}
				</span>
			</td>
			<td className="py-3 px-4 text-slate-600 dark:text-slate-400">
				{reason}
			</td>
			<td className="py-3 px-4 text-slate-500 dark:text-slate-500 text-sm">
				{new Date(createdAt).toLocaleDateString()}
			</td>
		</tr>
	);
}

export default Transaction;
