package accounts_routes

import (
	"backend/internal/controllers/controller_models"
	"backend/internal/models"
	"backend/internal/models/dto"
	"backend/internal/sdkhttp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccountByID(c *gin.Context) {
	// Parse request
	var req = dto.CreateAccountReqDto{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.CreateAccountResDto{
			BaseReply: controller_models.BaseReply{
				FailureCode: models.FailureCodeParseRequest,
				Error:       err.Error(),
			},
		})
		return
	}

	// Create account
	account := models.Account{
		AccountID: req.AccountId,
		Balance:   req.InitialBalance,
	}

	// Save account
	createdAccount, err := sdkhttp.Server.AccountsService.Create(c, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CreateAccountResDto{
			BaseReply: controller_models.BaseReply{
				FailureCode: models.FailureCodeServiceFailed,
				Error:       err.Error(),
			},
		})
		return
	}

	res := dto.CreateAccountResDto{
		AccountId: createdAccount.AccountID,
		Balance:   createdAccount.Balance,
	}

	c.JSON(http.StatusOK, res)
}

func GetAccountByID(c *gin.Context) {
	accountId := c.Param("account_id")

	// Get account
	account, err := sdkhttp.Server.AccountsService.GetByID(c, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GetAccountResDto{
			BaseReply: controller_models.BaseReply{
				FailureCode: models.FailureCodeServiceFailed,
				Error:       err.Error(),
			},
		})
		return
	}

	if account == nil {
		c.JSON(http.StatusNotFound, dto.GetAccountResDto{
			BaseReply: controller_models.BaseReply{
				FailureCode: models.FailureCodeNotFound,
				Error:       "account not found",
			},
		})
		return
	}

	res := dto.GetAccountResDto{
		AccountId: account.AccountID,
		Balance:   account.Balance,
	}

	c.JSON(http.StatusOK, res)
}