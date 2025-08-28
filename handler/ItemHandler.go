package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/service"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	Service *service.ItemService
}

func NewItemHandler(s *service.ItemService) *ItemHandler {
	return &ItemHandler{Service: s}
}

func (h *ItemHandler) Create(c *gin.Context) {
	clientCode := c.Query("client_code") // mandatory
	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateItem(&item, clientCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	clientCode := c.Query("client_code") // get client_code from query

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	item, err := h.Service.GetItem(id, clientCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) GetAll(c *gin.Context) {
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

	items, total, err := h.Service.GetItems(page, pageSize, sortBy, order, clientCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      items,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"sort":      sortBy,
		"order":     order,
		"totalPage": int((total + int64(pageSize) - 1) / int64(pageSize)), // ceil
	})
}

func (h *ItemHandler) SetActive(c *gin.Context) {
	idParam := c.Param("id")
	clientCode := c.Query("client_code") // string

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing client code"})
		return
	}

	if err := h.Service.SetActive(c.Request.Context(), id, clientCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item activated successfully"})
}

func (h *ItemHandler) SetInactive(c *gin.Context) {
	idParam := c.Param("id")
	clientCode := c.Query("client_code") // string

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing client code"})
		return
	}

	if err := h.Service.SetInactive(c.Request.Context(), id, clientCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item deactivated successfully"})
}

func (h *ItemHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	// get clientCode (example: from query param, adjust if you use header or body)
	clientCode := c.Query("client_code")
	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = id

	if err := h.Service.UpdateItem(&item, clientCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	clientCode := c.Query("client_code") // get client_code from query

	if clientCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_code is required"})
		return
	}

	if err := h.Service.DeleteItem(id, clientCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
}
