package app

import (
	"github.com/Barbarpotato/Storix/handler"
	"github.com/Barbarpotato/Storix/repository"
	"github.com/Barbarpotato/Storix/service"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB

	UnitRepo    *repository.UnitRepository
	UnitService *service.UnitService
	UnitHandler *handler.UnitHandler

	ClientRepo    *repository.ClientRepository
	ClientService *service.ClientService
	ClientHandler *handler.ClientHandler

	WarehouseRepo    *repository.WarehouseRepository
	WarehouseService *service.WarehouseService
	WarehouseHandler *handler.WarehouseHandler

	ItemRepo    *repository.ItemRepository
	ItemService *service.ItemService
	ItemHandler *handler.ItemHandler

	StockCardRepo    *repository.StockCardRepository
	StockCardService *service.StockCardService
	StockCardHandler *handler.StockCardHandler
}

func NewApp(db *gorm.DB) *App {
	unitRepo := repository.NewUnitRepository(db)
	unitService := service.NewUnitService(unitRepo)
	unitHandler := handler.NewUnitHandler(unitService)

	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(clientService)

	warehouseRepo := repository.NewWarehouseRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)

	itemRepo := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	stockCardRepo := repository.NewStockCardRepository(db)
	stockCardService := service.NewStockCardService(stockCardRepo)
	stockCardHandler := handler.NewStockCardHandler(stockCardService)

	return &App{
		DB:          db,
		UnitRepo:    unitRepo,
		UnitService: unitService,
		UnitHandler: unitHandler,

		ClientRepo:    clientRepo,
		ClientService: clientService,
		ClientHandler: clientHandler,

		WarehouseRepo:    warehouseRepo,
		WarehouseService: warehouseService,
		WarehouseHandler: warehouseHandler,

		ItemRepo:    itemRepo,
		ItemService: itemService,
		ItemHandler: itemHandler,

		StockCardRepo:    stockCardRepo,
		StockCardService: stockCardService,
		StockCardHandler: stockCardHandler,
	}
}
