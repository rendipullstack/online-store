package services

import (
	"online-store/helpers"
	"online-store/models"
	"online-store/repositories"
)

type CategoryServices interface {
	Create(input models.CategoryInput) (models.CategoryResponsePost, error)
	GetCategories() ([]models.CategoryResponseGet, error)
	SaveCategory(input models.CategoryInput, categoryID uint) (models.CategoryResponsePatch, error)
	DeleteCategory(categoryID uint) error
}

type categoryServices struct {
	repository repositories.CategoryRepositories
}

func NewCategoryServices(repository repositories.CategoryRepositories) *categoryServices {
	return &categoryServices{repository}
}

func (cs *categoryServices) Create(input models.CategoryInput) (models.CategoryResponsePost, error) {
	var (
		category         models.Category
		categoryResponse models.CategoryResponsePost
	)

	category.Type = input.Type

	category, err := cs.repository.PostCategory(category)

	categoryResponse.ID = category.ID
	categoryResponse.Type = category.Type
	categoryResponse.CreatedAt = category.CreatedAt

	return categoryResponse, helpers.ReturnIfError(err)
}

func (cs *categoryServices) GetCategories() ([]models.CategoryResponseGet, error) {
	var (
		categories          []models.Category
		categoriesResponses []models.CategoryResponseGet
	)

	categories, err := cs.repository.GetAllCategories()
	if err != nil {
		helpers.PanicIfError(err)
	}

	for _, category := range categories {
		var categoryResponse models.CategoryResponseGet

		categoryResponse.ID = category.ID
		categoryResponse.Type = category.Type
		categoryResponse.CreatedAt = category.CreatedAt
		categoryResponse.UpdatedAt = category.UpdatedAt
		var productResponses []models.ProductResponse
		for _, product := range category.Product {
			productResponse := models.ProductResponse{
				ID:         product.ID,
				Price:      product.Price,
				Title:      product.Title,
				Stock:      *product.Stock,
				CategoryID: product.CategoryID,
				CreatedAt:  product.CreatedAt,
			}
			productResponses = append(productResponses, productResponse)
		}
		categoryResponse.Product = productResponses

		categoriesResponses = append(categoriesResponses, categoryResponse)
	}

	return categoriesResponses, nil
}

func (cs *categoryServices) SaveCategory(input models.CategoryInput, categoryID uint) (models.CategoryResponsePatch, error) {
	var (
		category         models.Category
		categoryResponse models.CategoryResponsePatch
	)

	category, err := cs.repository.GetDataByID(categoryID)
	if err != nil {
		return categoryResponse, err
	}

	category.Type = input.Type

	category, err = cs.repository.UpdateCategory(category)

	categoryResponse.ID = category.ID
	categoryResponse.Type = category.Type
	categoryResponse.UpdatedAt = category.UpdatedAt

	return categoryResponse, helpers.ReturnIfError(err)

}

func (cs *categoryServices) DeleteCategory(categoryID uint) error {
	category, err := cs.repository.GetDataByID(categoryID)
	if err != nil {
		return err
	}

	err = cs.repository.DeleteCategory(category)

	return helpers.ReturnIfError(err)
}
