package models

type User struct {
    ID    int32  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateTransactionPayload struct {
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
}

type Transaction struct {
	ID        *int32 `json:"id"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"createdAt"`
}

const (
	TypeSaving     = "saving"
	TypeWithdrawal = "withdrawal"
)



type RegisterPayload struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type LoginPayload struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

