package controllers

import (
	"net/http"
	"online-store/models"
	"online-store/services"
	"online-store/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	s services.CartService
}

func NewCartController(s services.CartService) CartController {
	return CartController{s}
}

func (controller *CartController) CreateCart(c *gin.Context) {
	var input models.CartInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create cart",
				utils.GetErrorData(err),
			),
		)
		return
	}

	userID := c.GetUint("userID")

	cart, err := controller.s.CreateCart(input, userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create cart",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.NewResponse(
			http.StatusCreated,
			"Success to create cart",
			cart,
		),
	)
}

func (controller *CartController) GetUserCarts(c *gin.Context) {
	userID := c.GetUint("userID")

	carts, err := controller.s.GetCarts(userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get carts",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Success to get carts",
			carts,
		),
	)
}

func (controller *CartController) DeleteCart(c *gin.Context) {
	var cartIDRaw = c.Param("id")

	cartID, err := strconv.Atoi(cartIDRaw)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Parameter must be a valid ID",
				utils.GetErrorData(err),
			),
		)
		return
	}

	err = controller.s.DeleteCart(uint(cartID))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to delete data",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Successfully delete data",
			nil,
		),
	)
}
