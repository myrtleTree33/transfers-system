package transactions_routes

import (
	"backend/internal/controllers/controller_models"
	"backend/internal/models"
	"backend/internal/models/dto"
	"backend/internal/sdkhttp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func CreateTransaction(c *gin.Context) {
	// Parse request
	var req = dto.CreateTransactionReqDto{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.CreateTransactionResDto{
			BaseReply: controller_models.BaseReply{
				FailureCode: models.FailureCodeParseRequest,
				Error:       err.Error(),
			},
		})
		return
	}

	// Create transaction
	id := ksuid.New()
	transaction := models.Transaction{
		Base:                 models.Base{ID: &id},
		SourceAccountID:      req.SourceAccountID,
		DestinationAccountID: req.DestinationAccountID,
		Amount:               req.Amount,
	}

	// Save transaction
	createdTransaction, err := sdkhttp.Server.TransactionsService.Create(c, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CreateTransactionResDto{
			BaseReply: controller_models.BaseReply{
				FailureCode: models.FailureCodeServiceFailed,
				Error:       err.Error(),
			},
		})
		return
	}

	res := dto.CreateTransactionResDto{
		Transaction: *createdTransaction,
	}

	c.JSON(http.StatusOK, res)
}
