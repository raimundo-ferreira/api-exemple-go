package main

import (
	"api-exemple/pkg/configs"
	"api-exemple/pkg/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// @Title API with Golang
// @Version	1.0
// @Description	Example API with Golang using Gin, PostgreSQL, PGX, JWT and Swagger.
// @Contact.name API Support
// @Contact.email byraimundoferreira@gmail.com
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @BasePath /api
func main() {
	configs.LoadEnvVariables()

	db, err := configs.OpenDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.Initialize(router, db)

	if err := router.Run(os.Getenv("SERVER_URL")); err != nil {
		log.Fatalf("Failed... Server is not running! Detalhes: %v", err)
	}
}
