package routes

import (
	"api-exemple/app/controller"
	"api-exemple/app/repository"
	"api-exemple/app/service"
	"api-exemple/pkg/middleware"

	_ "api-exemple/docs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initialize(router *gin.Engine, db *pgx.Conn) {
	// Repository
	userRepository := repository.NewUserRepositoryImpl(db)
	productRepository := repository.NewProductRepositoryImpl(db)

	// Service
	userService := service.NewUserServiceImpl(userRepository)
	productService := service.NewProductServiceImpl(productRepository)

	// Controller
	authController := controller.NewAuthController(userService)
	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)

	// Eneble CORS
	router.Use(middleware.CORSMiddleware())

	//public routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/register", authController.Register)
	router.POST("/api/token", authController.Token)

	//private routes
	router.Use(middleware.JwtAuthMiddleware())
	router.GET("/api/state", authController.State)
	UserRegistreRoute(router, userController)
	ProductRegistreRoute(router, productController)
}
