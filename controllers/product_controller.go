package controllers

import (
	"net/http"
	"online-store/models"
	"online-store/services"
	"online-store/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductServices
}

func NewProductController(productService services.ProductServices) ProductController {
	return ProductController{productService}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input models.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Something wrong with input",
				utils.GetErrorData(err),
			),
		)
		return
	}

	product, err := pc.productService.Create(input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create product",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.NewResponse(
			http.StatusCreated,
			"Successfuly create product",
			product,
		),
	)
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.productService.GetProducts()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get data",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Successfully get data",
			products,
		),
	)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	var productIDRaw = c.Param("id")

	productID, err := strconv.Atoi(productIDRaw)
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

	err = pc.productService.DeleteProduct(uint(productID))
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
			"Successfuly delete data",
			nil,
		),
	)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var (
		input        models.ProductInput
		productIDRaw = c.Param("id")
	)
	productID, err := strconv.Atoi(productIDRaw)
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

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Something wrong with input",
				utils.GetErrorData(err),
			),
		)
		return
	}

	product, err := pc.productService.SaveProduct(input, uint(productID))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update data",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Successfully update data",
			product,
		),
	)
}
