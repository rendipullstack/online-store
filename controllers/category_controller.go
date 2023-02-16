package controllers

import (
	"net/http"
	"online-store/models"
	"online-store/services"
	"online-store/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.CategoryServices
}

func NewCategoryController(service services.CategoryServices) CategoryController {
	return CategoryController{service}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var input models.CategoryInput
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

	category, err := cc.categoryService.Create(input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create category",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.NewResponse(
			http.StatusCreated,
			"Created successfuly",
			category,
		),
	)
}

func (cc *CategoryController) GetCategories(c *gin.Context) {
	categories, err := cc.categoryService.GetCategories()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all data",
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
			categories,
		),
	)
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var (
		input         models.CategoryInput
		categoryIDRaw = c.Param("id")
	)

	categoryID, err := strconv.Atoi(categoryIDRaw)
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

	category, err := cc.categoryService.SaveCategory(input, uint(categoryID))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to patch data",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Successfuly patch data",
			category,
		),
	)

}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	var categoryIDRaw = c.Param("id")

	categoryID, err := strconv.Atoi(categoryIDRaw)
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

	err = cc.categoryService.DeleteCategory(uint(categoryID))
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
