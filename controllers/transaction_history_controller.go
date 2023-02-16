package controllers

import (
	"net/http"
	"online-store/models"
	"online-store/services"
	"online-store/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	s services.TransactionService
}

func NewTransactionController(s services.TransactionService) TransactionController {
	return TransactionController{s}
}

func (controller *TransactionController) CreateTransaction(c *gin.Context) {
	var input models.TransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create transaction",
				utils.GetErrorData(err),
			),
		)
		return
	}

	userID := c.GetUint("userID")

	transaction, err := controller.s.CreateTransaction(input, userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create transaction",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.NewResponse(
			http.StatusCreated,
			"Success to create transaction",
			transaction,
		),
	)
}

func (controller *TransactionController) GetUserTransactions(c *gin.Context) {
	userID := c.GetUint("userID")

	transactions, err := controller.s.GetTransactions(userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get transactions",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Success to get transactions",
			transactions,
		),
	)
}

func (controller *TransactionController) GetAllTransactions(c *gin.Context) {
	transactions, err := controller.s.GetAllTransactions()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all transactions",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Success to get transactions",
			transactions,
		),
	)
}
