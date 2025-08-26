package handler

import (
	"net/http"
	"strconv"

	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/service"

	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	Service *service.WarehouseService
}

func NewWarehouseHandler(s *service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{Service: s}
}

func (h *WarehouseHandler) Create(c *gin.Context) {
	var warehouse models.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateWarehouse(&warehouse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, warehouse)
}

func (h *WarehouseHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	warehouse, err := h.Service.GetWarehouse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "warehouse not found"})
		return
	}
	c.JSON(http.StatusOK, warehouse)
}

func (h *WarehouseHandler) GetAll(c *gin.Context) {
	warehouses, err := h.Service.GetWarehouses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, warehouses)
}

func (h *WarehouseHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var warehouse models.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	warehouse.ID = id
	if err := h.Service.UpdateWarehouse(&warehouse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, warehouse)
}

func (h *WarehouseHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.Service.DeleteWarehouse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "warehouse deleted"})
}
