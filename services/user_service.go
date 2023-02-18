package services

import (
	"errors"
	_ "online-store/docs"
	"online-store/helpers"
	"online-store/models"
	"online-store/repositories"
)

type UserService interface {
	Register(input models.UserRegisterInput) (models.UserResponseRegister, error)
	Login(input models.UserLoginInput) (string, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

// Register godoc
// @Summary      Register account
// @Description  Register an account
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        request body models.UserRegisterInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/register [post]
func (us *userService) Register(input models.UserRegisterInput) (models.UserResponseRegister, error) {
	var (
		user         models.User
		userResponse models.UserResponseRegister
	)
	user, _ = us.repository.GetUserByEmail(input.Email)
	if user.ID > 0 {
		return userResponse, errors.New("email already existed")
	}

	// Hash Password Here
	password, err := helpers.HashPassword(input.Password)
	if err != nil {
		return userResponse, errors.New("something wrong with password")
	}

	user.FullName = input.FullName
	user.Email = input.Email
	user.Password = password
	user.Role = "customer"

	if user.Role == "admin" {
		return userResponse, errors.New("failed to upload your data")
	}

	user, err = us.repository.CreateUser(user)

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.Password = user.Password
	userResponse.CreatedAt = user.CreatedAt

	return userResponse, helpers.ReturnIfError(err)
}

// Login godoc
// @Summary      Login account
// @Description  Login an account
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        request body models.UserLoginInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/login [post]
func (us *userService) Login(input models.UserLoginInput) (string, error) {
	var token string

	user, _ := us.repository.GetUserByEmail(input.Email)
	if user.ID == 0 {
		return "", errors.New("user is not existed")
	}

	ok := helpers.ComparePassword(user.Password, input.Password)
	if !ok {
		return token, errors.New("password is wrong")
	}

	token, err := helpers.GenerateToken(user)
	return token, helpers.ReturnIfError(err)

}
