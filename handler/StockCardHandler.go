package handler

import (
	"net/http"
	"strconv"

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
	var stockCard models.StockCard
	if err := c.ShouldBindJSON(&stockCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateStockCard(&stockCard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, stockCard)
}

func (h *StockCardHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	stockCard, err := h.Service.GetStockCard(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "stock card not found"})
		return
	}
	c.JSON(http.StatusOK, stockCard)
}

func (h *StockCardHandler) GetAll(c *gin.Context) {
	stockCards, err := h.Service.GetStockCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stockCards)
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
	if err := h.Service.DeleteStockCard(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "stock card deleted"})
}
