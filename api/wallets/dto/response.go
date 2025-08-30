package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type WalletRes struct {
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type WalletTransactionRes struct {
	Amount    int64     `json:"amount"`
	Type      string    `json:"type"`
	Reference string    `json:"reference"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

type WalletPayoutRes struct {
	UUID            string    `json:"uuid"`
	Amount          int64     `json:"amount"`
	Status          string    `json:"status"`
	BankName        string    `json:"bank_name"`
	BankAccount     string    `json:"bank_account"`
	BankAccountName string    `json:"bank_account_name"`
	CreatedAt       time.Time `json:"created_at"`

	TransactionReceipt WalletPayoutTransactionReceiptRes `json:"transaction_receipt,omitempty"`
	User               UserRes                           `json:"user"`
}

type WalletPayoutTransactionReceiptRes struct {
	ReceiptURL string `json:"receipt_url,omitempty"`
	Note       string `json:"note,omitempty"`
}

type UserRes struct {
	UUID            string `json:"uuid"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	AvatarURL       string `json:"avatar_url"`
}
