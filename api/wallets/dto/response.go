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