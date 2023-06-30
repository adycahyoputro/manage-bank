package dto

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EntryRequest struct {
	Amount int64 `json:"amount"`
}

type TransferRequest struct {
	ToAccountID string `json:"to_account_id"`
	Amount      int64  `json:"amount"`
}

type AccountRequest struct {
	AccountID string `json:"account_id"`
	Balance   int64  `json:"balance"`
}

type LogoutRequest struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}
