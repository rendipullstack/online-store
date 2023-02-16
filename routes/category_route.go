package routes

import (
	"online-store/controllers"
	"online-store/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, controller controllers.CategoryController) {
	router.POST("/categories", middlewares.AuthMiddleware, middlewares.AdminMiddleware, controller.CreateCategory)
	router.GET("/categories", middlewares.AuthMiddleware, controller.GetCategories)
	router.PATCH("/categories/:id", middlewares.AuthMiddleware, middlewares.AdminMiddleware, controller.UpdateCategory)
	router.DELETE("/categories/:id", middlewares.AuthMiddleware, middlewares.AdminMiddleware, controller.DeleteCategory)
}
