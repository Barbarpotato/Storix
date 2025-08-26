package handler

import (
	"net/http"
	"strconv"

	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/service"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	Service *service.ClientService
}

func NewClientHandler(s *service.ClientService) *ClientHandler {
	return &ClientHandler{Service: s}
}

func (h *ClientHandler) Create(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	client, err := h.Service.GetClient(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) GetAll(c *gin.Context) {
	clients, err := h.Service.GetClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (h *ClientHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client.ID = id
	if err := h.Service.UpdateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.Service.DeleteClient(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "client deleted"})
}
