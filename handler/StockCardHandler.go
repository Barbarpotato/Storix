package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/service"

	"github.com/gin-gonic/gin"
)

type StockCardHandler struct {
	Service *service.StockCardService
}

func NewStockCardHandler(s *service.StockCardService) *StockCardHandler {
	return &StockCardHandler{Service: s}
}

func (h *StockCardHandler) Create(c *gin.Context) {
	clientCode := c.Query("client_code") // mandatory
	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	var stockCard models.StockCard
	if err := c.ShouldBindJSON(&stockCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateStockCard(&stockCard, clientCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, stockCard)
}

func (h *StockCardHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	clientCode := c.Query("client_code") // get client_code from query

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	stockCard, err := h.Service.GetStockCard(id, clientCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "stock card not found"})
		return
	}
	c.JSON(http.StatusOK, stockCard)
}

func (h *StockCardHandler) GetAll(c *gin.Context) {
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

	stockCards, total, err := h.Service.GetStockCards(page, pageSize, sortBy, order, clientCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      stockCards,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"sort":      sortBy,
		"order":     order,
		"totalPage": int((total + int64(pageSize) - 1) / int64(pageSize)), // ceil
	})
}

func (h *StockCardHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var stockCard models.StockCard
	if err := c.ShouldBindJSON(&stockCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stockCard.ID = id
	if err := h.Service.UpdateStockCard(&stockCard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stockCard)
}

func (h *StockCardHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	clientCode := c.Query("client_code") // get client_code from query

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	if err := h.Service.DeleteStockCard(id, clientCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "stock card deleted"})
}
