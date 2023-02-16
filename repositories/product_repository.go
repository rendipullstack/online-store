package repositories

import (
	"online-store/helpers"
	"online-store/models"

	"gorm.io/gorm"
)

type ProductRepositories interface {
	PostProduct(product models.Product) (models.Product, error)
	GetProducts() ([]models.Product, error)
	DeleteProducts(product models.Product) error
	GetDataByID(productID uint) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
}

type productRepositories struct {
	db *gorm.DB
}

func NewProductRepositories(db *gorm.DB) *productRepositories {
	return &productRepositories{db}
}

func (pr *productRepositories) PostProduct(product models.Product) (models.Product, error) {
	err := pr.db.Create(&product).Error
	return product, helpers.ReturnIfError(err)
}

func (pr *productRepositories) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := pr.db.Find(&products).Error
	return products, helpers.ReturnIfError(err)
}

func (pr *productRepositories) DeleteProducts(product models.Product) error {
	err := pr.db.Delete(&product).Error
	return helpers.ReturnIfError(err)
}

func (pr *productRepositories) GetDataByID(productID uint) (models.Product, error) {
	var product models.Product
	err := pr.db.Preload("Category").Where("id = ?", productID).Find(&product).Error
	return product, helpers.ReturnIfError(err)
}

func (pr *productRepositories) UpdateProduct(product models.Product) (models.Product, error) {
	err := pr.db.Where("id = ?", product.ID).Updates(&product).Error
	return product, helpers.ReturnIfError(err)
}
