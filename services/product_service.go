package services

import (
	"fmt"
	"online-store/helpers"
	"online-store/models"
	"online-store/repositories"

	"github.com/dustin/go-humanize"
)

type ProductServices interface {
	Create(input models.ProductInput) (models.ProductResponse, error)
	GetProducts() ([]models.ProductResponse, error)
	DeleteProduct(productID uint) error
	SaveProduct(input models.ProductInput, productID uint) (models.ProductResponseUpdate, error)
}

type productServices struct {
	repository         repositories.ProductRepositories
	categoryRepository repositories.CategoryRepositories
}

func NewProductServices(repository repositories.ProductRepositories, categoryRepository repositories.CategoryRepositories) *productServices {
	return &productServices{repository, categoryRepository}
}

// Create Product godoc
// @Summary      Create product
// @Description  Add a new product role only admin
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        request body models.ProductInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /products [post]
// @Security BearerAuth
func (ps *productServices) Create(input models.ProductInput) (models.ProductResponse, error) {
	var (
		product         models.Product
		productResponse models.ProductResponse
	)

	category, err := ps.categoryRepository.GetDataByID(input.CategoryID)
	if err != nil {
		return productResponse, err
	}

	product.Title = input.Title
	product.Price = input.Price
	product.Stock = &input.Stock
	product.CategoryID = category.ID

	product, err = ps.repository.PostProduct(product)

	productResponse.ID = product.ID
	productResponse.Title = product.Title
	productResponse.Price = product.Price
	productResponse.Stock = *product.Stock
	productResponse.CategoryID = product.CategoryID
	productResponse.CreatedAt = product.CreatedAt

	return productResponse, helpers.ReturnIfError(err)
}

// Product details godoc
// @Summary      Product details
// @Description  Product details
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /products [get]
// @Security BearerAuth
func (ps *productServices) GetProducts() ([]models.ProductResponse, error) {
	var (
		products         []models.Product
		productResponses []models.ProductResponse
	)

	products, err := ps.repository.GetProducts()

	for _, product := range products {
		var productResponse models.ProductResponse

		productResponse.ID = product.ID
		productResponse.Title = product.Title
		productResponse.Price = product.Price
		productResponse.Stock = *product.Stock
		productResponse.CategoryID = product.CategoryID
		productResponse.CreatedAt = product.CreatedAt

		productResponses = append(productResponses, productResponse)
	}

	return productResponses, helpers.ReturnIfError(err)
}

// Delete Product godoc
// @Summary      Delete product
// @Description  Delete product
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /products/1 [delete]
// @Security BearerAuth
func (ps *productServices) DeleteProduct(productID uint) error {
	product, err := ps.repository.GetDataByID(productID)
	if err != nil {
		return err
	}

	err = ps.repository.DeleteProducts(product)

	return helpers.ReturnIfError(err)
}

// Update Product godoc
// @Summary      Update product
// @Description  Update product
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        request body models.ProductInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /products/1 [put]
// @Security BearerAuth
func (ps *productServices) SaveProduct(input models.ProductInput, productID uint) (models.ProductResponseUpdate, error) {
	var (
		product         models.Product
		productResponse models.ProductResponseUpdate
	)

	product, err := ps.repository.GetDataByID(productID)
	if err != nil {
		return productResponse, err
	}

	product.Title = input.Title
	product.Price = input.Price
	product.Stock = &input.Stock
	product.CategoryID = input.CategoryID

	product, err = ps.repository.UpdateProduct(product)

	productData := models.ProductResponseData{
		ID:         productID,
		Title:      product.Title,
		Price:      fmt.Sprintf("Rp %s", humanize.Comma(int64(product.Price))),
		Stock:      *product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}

	productResponse = models.ProductResponseUpdate{
		Product: productData,
	}

	return productResponse, helpers.ReturnIfError(err)
}
