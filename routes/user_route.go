package routes

import (
	"online-store/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, controller controllers.UserController) {
	router.POST("users/register", controller.SignUp)
	router.POST("users/login", controller.SignIn)
}
