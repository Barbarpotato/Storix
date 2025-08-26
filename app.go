package main

import (
	"github.com/Barbarpotato/Storix/handler"
	"github.com/Barbarpotato/Storix/repository"
	"github.com/Barbarpotato/Storix/service"
	"gorm.io/gorm"
)

type App struct {

	DB *gorm.DB

	UnitRepo     *repository.UnitRepository
	UnitService  *service.UnitService
	UnitHandler  *handler.UnitHandler

	// add more as needed
}

func NewApp(db *gorm.DB) *App {
	unitRepo := repository.NewUnitRepository(db)
	unitService := service.NewUnitService(unitRepo)
	unitHandler := handler.NewUnitHandler(unitService)

	return &App{
		DB:           db,
		UnitRepo:     unitRepo,
		UnitService:  unitService,
		UnitHandler:  unitHandler,
	}
}
