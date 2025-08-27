package mapper

import (
	"ragamaya-api/api/wallets/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapWalletMTO(input models.Wallet) dto.WalletRes {
	var output dto.WalletRes

	mapstructure.Decode(input, &output)
	return output
}

func MapWalletTransactionMTO(input models.WalletTransactionHistory) dto.WalletTransactionRes {
	var output dto.WalletTransactionRes

	mapstructure.Decode(input, &output)
	output.CreatedAt = input.CreatedAt
	return output
}
