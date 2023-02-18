package services

import (
	"errors"
	"online-store/helpers"
	"online-store/models"
	"online-store/repositories"
)

type CartService interface {
	CreateCart(input models.CartInput, userID uint) (models.CartPostResponse, error)
	GetCarts(userID uint) ([]models.UserCartResponse, error)
	DeleteCart(productID uint) error
}

type cartService struct {
	cartRepository     repositories.CartRepository
	productRepository  repositories.ProductRepositories
	userRepository     repositories.UserRepository
	categoryRepository repositories.CategoryRepositories
}

func NewCartService(
	cartRepository repositories.CartRepository,
	productRepository repositories.ProductRepositories,
	userRepository repositories.UserRepository,
	categoryRepository repositories.CategoryRepositories,
) *cartService {
	return &cartService{
		cartRepository,
		productRepository,
		userRepository,
		categoryRepository,
	}
}

// CreateCart godoc
// @Summary      Add to cart
// @Description  Add product to shopping cart
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        request body models.CartInput true "Payload Body [RAW]"
// @Success      200 {object} models.CartPostResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /cart [post]
// @Security BearerAuth
func (s *cartService) CreateCart(input models.CartInput, userID uint) (models.CartPostResponse, error) {
	var (
		cartResponse models.CartPostResponse
		cart         models.Cart
	)

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return cartResponse, err
	}
	product, err := s.productRepository.GetDataByID(uint(input.ProductID))
	if err != nil {
		return cartResponse, err
	}
	category, err := s.categoryRepository.GetDataByID(product.CategoryID)
	if err != nil {
		return cartResponse, err
	}

	if *product.Stock < input.Quantity {
		return cartResponse, errors.New("product is not available")
	}

	category, err = s.categoryRepository.UpdateCategory(category)
	if err != nil {
		return cartResponse, err
	}

	stock := *product.Stock - input.Quantity
	*product.Stock = stock
	product, err = s.productRepository.UpdateProduct(product)
	if err != nil {
		return cartResponse, err
	}

	cart.UserID = user.ID
	cart.ProductID = product.ID
	cart.Quantity = input.Quantity

	cart, err = s.cartRepository.CreateCart(cart)
	cartResponse = models.CartPostResponse{
		Quantity:     cart.Quantity,
		ProductTitle: product.Title,
	}

	return cartResponse, helpers.ReturnIfError(err)
}

// GetCarts godoc
// @Summary      Cart
// @Description  List of products that have been added to the shopping cart
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Success      200 {object} models.UserCartResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /cart [get]
// @Security BearerAuth
func (s *cartService) GetCarts(userID uint) ([]models.UserCartResponse, error) {
	carts, err := s.cartRepository.GetCarts(userID)
	var cartResponses []models.UserCartResponse
	for _, cart := range carts {
		cartResponse := models.UserCartResponse{
			ID:        cart.ID,
			ProductID: cart.ProductID,
			UserID:    cart.UserID,
			Quantity:  cart.Quantity,
			Product: models.ProductResponse{
				ID:         cart.Product.ID,
				Title:      cart.Product.Title,
				Price:      cart.Product.Price,
				Stock:      *cart.Product.Stock,
				CategoryID: cart.Product.CategoryID,
				CreatedAt:  cart.Product.CreatedAt,
			},
		}
		cartResponses = append(cartResponses, cartResponse)
	}

	return cartResponses, helpers.ReturnIfError(err)
}

// GetCarts godoc
// @Summary      Delete cart
// @Description  Delete product list in shopping cart
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /cart/1 [delete]
// @Security BearerAuth
func (s *cartService) DeleteCart(cartID uint) error {

	cart, err := s.cartRepository.GetDataByID(cartID)
	if err != nil {
		return err
	}

	product, err := s.productRepository.GetDataByID(uint(cart.ProductID))
	if err != nil {
		return err
	}

	stock := *product.Stock + cart.Quantity
	*product.Stock = stock
	product, err = s.productRepository.UpdateProduct(product)
	if err != nil {
		return err
	}

	err = s.cartRepository.DeleteCart(cart)

	return helpers.ReturnIfError(err)
}
