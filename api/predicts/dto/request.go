package dto

type PredictReq struct {
	File []byte `validate:"required"`
}
