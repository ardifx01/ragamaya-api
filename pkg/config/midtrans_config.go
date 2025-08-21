package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitMidtrans() *coreapi.Client {
	midtransServerKey := os.Getenv("MIDTRANS_SERVER_KEY")
	midtransCore := coreapi.Client{}

	if os.Getenv("MIDTRANS_ENV") == "sandbox" {
		midtransCore.New(midtransServerKey, midtrans.Sandbox)
	} else if os.Getenv("MIDTRANS_ENV") == "production" {
		midtransCore.New(midtransServerKey, midtrans.Production)
	}

	return &midtransCore
}
