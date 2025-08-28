package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/service"

	"github.com/gin-gonic/gin"
)

type UnitHandler struct {
	Service *service.UnitService
}

func NewUnitHandler(s *service.UnitService) *UnitHandler {
	return &UnitHandler{Service: s}
}

func (h *UnitHandler) Create(c *gin.Context) {
	var unit models.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateUnit(&unit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, unit)
}

func (h *UnitHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	unit, err := h.Service.GetUnit(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "unit not found"})
		return
	}
	c.JSON(http.StatusOK, unit)
}

func (h *UnitHandler) GetAll(c *gin.Context) {
	// defaults
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	sortBy := c.DefaultQuery("sort", "id")
	order := strings.ToLower(c.DefaultQuery("order", "asc"))

	units, total, err := h.Service.GetUnits(page, pageSize, sortBy, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      units,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"sort":      sortBy,
		"order":     order,
		"totalPage": int((total + int64(pageSize) - 1) / int64(pageSize)), // ceil
	})
}

func (h *UnitHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var unit models.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unit.ID = id
	if err := h.Service.UpdateUnit(&unit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, unit)
}

func (h *UnitHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.Service.DeleteUnit(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "unit deleted"})
}
