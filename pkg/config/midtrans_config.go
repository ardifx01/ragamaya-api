package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitMidtrans() *coreapi.Client {
	midtransServerKey := GetMidtransServerKey()
	midtransCore := coreapi.Client{}

	if IsMidtransProduction() {
		midtransCore.New(midtransServerKey, midtrans.Production)
	} else {
		midtransCore.New(midtransServerKey, midtrans.Sandbox)
	}

	return &midtransCore
}
