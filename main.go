package main

import (
	"github.com/Barbarpotato/Storix/app"
	"github.com/Barbarpotato/Storix/configs"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to DB and capture the *gorm.DB instance
	configs.ConnectDB()

	app := app.NewApp(configs.DB)

	r := gin.Default()

	unitRoutes := r.Group("/units")
	{
		unitRoutes.POST("/", app.UnitHandler.Create)
		unitRoutes.GET("/", app.UnitHandler.GetAll)
		unitRoutes.GET("/:id", app.UnitHandler.Get)
		unitRoutes.PUT("/:id", app.UnitHandler.Update)
		unitRoutes.DELETE("/:id", app.UnitHandler.Delete)
	}

	clientRoutes := r.Group("/clients")
	{
		clientRoutes.POST("/", app.ClientHandler.Create)
		clientRoutes.GET("/", app.ClientHandler.GetAll)
		clientRoutes.GET("/:id", app.ClientHandler.Get)
		clientRoutes.PUT("/:id", app.ClientHandler.Update)
		clientRoutes.DELETE("/:id", app.ClientHandler.Delete)
	}

	warehouseRoutes := r.Group("/warehouses")
	{
		warehouseRoutes.POST("/", app.WarehouseHandler.Create)
		warehouseRoutes.GET("/", app.WarehouseHandler.GetAll)
		warehouseRoutes.GET("/:id", app.WarehouseHandler.Get)
		warehouseRoutes.PUT("/:id", app.WarehouseHandler.Update)
		warehouseRoutes.DELETE("/:id", app.WarehouseHandler.Delete)
	}

	itemRoutes := r.Group("/items")
	{
		itemRoutes.POST("/", app.ItemHandler.Create)
		itemRoutes.GET("/", app.ItemHandler.GetAll)
		itemRoutes.GET("/:id", app.ItemHandler.Get)
		itemRoutes.PUT("/:id", app.ItemHandler.Update)
		itemRoutes.DELETE("/:id", app.ItemHandler.Delete)
	}

	stockCardRoutes := r.Group("/stock-cards")
	{
		stockCardRoutes.POST("/", app.StockCardHandler.Create)
		stockCardRoutes.GET("/", app.StockCardHandler.GetAll)
		stockCardRoutes.GET("/:id", app.StockCardHandler.Get)
		stockCardRoutes.PUT("/:id", app.StockCardHandler.Update)
		stockCardRoutes.DELETE("/:id", app.StockCardHandler.Delete)
	}

	r.Run(":8080")
}
