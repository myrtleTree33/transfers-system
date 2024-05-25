package dto

import "backend/internal/controllers/controller_models"

type CreateAccountReqDto struct {
	AccountId      string  `json:"account_id" binding:"required"`
	InitialBalance float64 `json:"initial_balance" binding:"required"`
}

type CreateAccountResDto struct {
	controller_models.BaseReply
	AccountId string  `json:"account_id" binding:"required"`
	Balance   float64 `json:"balance" binding:"required"`
}
