export interface TransactionType {
  amount: number | string;
  reason: string;
  createdAt: string;
  type: "saving" | "withdrawal";
  id?: string;
}

export interface GetTransactionsParamsType {
  type?: "saving" | "withdrawal";
  size?: number;
}

export interface TransactionsResponse {
  transactions: TransactionType[];
  total_savings: string;
  total_withdrawals: string;
}