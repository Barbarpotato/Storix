package handler

import (
	"net/http"
	"strconv"
	"strings"

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
	clientCode := c.Query("client_code") // get client_code from query

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	warehouse, err := h.Service.GetWarehouse(id, clientCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "warehouse not found"})
		return
	}
	c.JSON(http.StatusOK, warehouse)
}

func (h *WarehouseHandler) GetAll(c *gin.Context) {
	clientCode := c.Query("client_code") // mandatory

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	// defaults
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	sortBy := c.DefaultQuery("sort", "name")
	order := strings.ToLower(c.DefaultQuery("order", "asc"))

	warehouses, total, err := h.Service.GetWarehouses(page, pageSize, sortBy, order, clientCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      warehouses,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"sort":      sortBy,
		"order":     order,
		"totalPage": int((total + int64(pageSize) - 1) / int64(pageSize)), // ceil
	})
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
	clientCode := c.Query("client_code") // get client_code from query

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	if err := h.Service.DeleteWarehouse(id, clientCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "warehouse deleted"})
}
