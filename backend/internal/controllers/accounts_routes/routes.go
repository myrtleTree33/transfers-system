package accounts_routes

import (
	"backend/internal/controllers/controller_models"
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccountReqDto struct {
	AccountId      string  `json:"account_id" binding:"required"`
	InitialBalance float64 `json:"initial_balance" binding:"required"`
}

type CreateAccountResDto struct {
	controller_models.BaseReply
	AccountId string  `json:"account_id" binding:"required"`
	Balance   float64 `json:"balance" binding:"required"`
}

func CreateAccountByID(c *gin.Context) {
	// Parse request
	var req = CreateAccountReqDto{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CreateAccountResDto{
			BaseReply: controller_models.BaseReply{
				FailureCode: models.FailureCodeParseRequest,
				Error:       err.Error(),
			},
		})
		return
	}

	res := CreateAccountResDto{
		AccountId: req.AccountId,
		Balance:   1000.00,
	}

	c.JSON(http.StatusOK, res)
}
