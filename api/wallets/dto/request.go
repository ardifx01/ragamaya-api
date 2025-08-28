package dto

type WalletTransactionReq struct {
	UserUUID  string `json:"user_uuid" validate:"required,uuid4"`
	Amount    int64  `json:"amount" validate:"required,gte=0"`
	Type      string `json:"type" validate:"required,oneof=debit credit"`
	Reference string `json:"reference" validate:"required"`
	Note      string `json:"note" validate:"required"`
}

type WalletPayoutReq struct {
	Amount          int64        `json:"amount" validate:"required,gte=50000"`
	BankName        string       `json:"bank_name" validate:"required,oneof=bca bni bri mandiri"`
	BankAccount     string       `json:"bank_account" validate:"required,number"`
	BankAccountName string       `json:"bank_account_name" validate:"required"`
}
