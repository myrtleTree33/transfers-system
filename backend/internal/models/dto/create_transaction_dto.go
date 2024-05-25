package dto

import (
	"backend/internal/controllers/controller_models"
	"backend/internal/models"
)

type CreateTransactionReqDto struct {
	SourceAccountID      string  `json:"source_account_id" binding:"required"`
	DestinationAccountID string  `json:"destination_account_id" binding:"required"`
	Amount               float64 `json:"amount" binding:"required"`
}

type CreateTransactionResDto struct {
	controller_models.BaseReply
	models.Transaction
}
