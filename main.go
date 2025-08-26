package main

import (
	"github.com/Barbarpotato/Storix/configs"
	"github.com/Barbarpotato/Storix/handler"
	"github.com/Barbarpotato/Storix/repository"
	"github.com/Barbarpotato/Storix/service"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to DB and capture the *gorm.DB instance
	configs.ConnectDB()

	r := gin.Default()

	// your existing endpoints remain
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Welcome to Gin 🚀")
	// })

	// r.GET("/data", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  "ok",
	// 		"message": "Hello from Gin",
	// 	})
	// })

	unitRepo := repository.NewUnitRepository(configs.DB)
	unitService := service.NewUnitService(unitRepo)
	unitHandler := handler.NewUnitHandler(unitService)

	unitRoutes := r.Group("/units")
	{
		unitRoutes.POST("/", unitHandler.Create)
		unitRoutes.GET("/", unitHandler.GetAll)
		unitRoutes.GET("/:id", unitHandler.Get)
		unitRoutes.PUT("/:id", unitHandler.Update)
		unitRoutes.DELETE("/:id", unitHandler.Delete)
	}

	r.Run(":8080")
}
