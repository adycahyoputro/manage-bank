package dto

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type UserResponse struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Balance  int64  `json:"balance"`
}

type TransferResponse struct {
	ID            string `json:"id"`
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`
	Amount        int64  `json:"amount"`
}

type EntryResponse struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
	Balance   int64  `json:"balance"`
}
