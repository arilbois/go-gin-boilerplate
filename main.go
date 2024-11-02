// main.go
package main

import (
	"msvc-syahril-app/config"
	_ "msvc-syahril-app/docs"
	"msvc-syahril-app/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title msvc-syahril-app API
// @version 1.0
// @description Ini adalah API untuk msvc-syahril-app.
// @host localhost:8080
// @BasePath /
func main() {
    cfg := config.GetConfig()
    config.ConnectDatabase(cfg)

    router := routes.SetupRouter()

    // Rute Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Menjalankan server
    router.Run(":" + cfg.AppPort)
}
