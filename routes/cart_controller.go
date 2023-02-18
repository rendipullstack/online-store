package routes

import (
	"online-store/controllers"
	"online-store/middlewares"

	"github.com/gin-gonic/gin"
)

func CartRoute(c *gin.Engine, controller controllers.CartController) {
	c.POST("/cart", middlewares.AuthMiddleware, controller.CreateCart)
	c.GET("/cart", middlewares.AuthMiddleware, controller.GetUserCarts)
	c.DELETE("/cart/:id", middlewares.AuthMiddleware, controller.DeleteCart)
}
