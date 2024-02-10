package routes

import (
	"api-exemple/app/controller"

	"github.com/gin-gonic/gin"
)

func UserRegistreRoute(router *gin.Engine, controller *controller.UserController) {
	v1 := router.Group("/api")
	{
		v1.GET("/user", controller.GetAll)
		v1.GET("/user/:id", controller.GetById)
		v1.POST("user", controller.Create)
		v1.PUT("user", controller.Update)
		v1.DELETE("user/:id", controller.Delete)
	}
}

func ProductRegistreRoute(router *gin.Engine, controller *controller.ProductController) {
	v1 := router.Group("/api")
	{
		v1.GET("/product", controller.GetAll)
		v1.GET("/product/:id", controller.GetById)
		v1.POST("product", controller.Create)
		v1.PUT("product", controller.Update)
		v1.DELETE("product/:id", controller.Delete)
	}
}
