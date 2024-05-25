package dto

import "backend/internal/controllers/controller_models"

type GetAccountResDto struct {
	controller_models.BaseReply
	AccountId string  `json:"account_id" binding:"required"`
	Balance   float64 `json:"balance" binding:"required"`
}
