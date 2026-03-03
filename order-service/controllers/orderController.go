package controllers

import (
	"net/http"
	"order-service/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var orders []models.OrderResponse
var idCounter int = 1

func GetOrders(c *gin.Context) {
	if orders == nil {
		orders = make([]models.OrderResponse, 0)
	}
	c.JSON(http.StatusOK, orders)
}

func PlaceOrder(c *gin.Context) {
	var orderReq models.OrderRequest

	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order data: " + err.Error(),
		})
		return
	}

	orderResp := models.OrderResponse{
		ID:         idCounter,
		CustomerID: orderReq.CustomerID,
		Items:      orderReq.Items,
		Total:      orderReq.Total,
		Address:    orderReq.Address,
		Status:     models.StatusPending,
		CreatedAt:  time.Now().Format(time.RFC3339),
		Metadata:   orderReq.Metadata,
	}

	idCounter++

	if orders == nil {
		orders = make([]models.OrderResponse, 0)
	}
	orders = append(orders, orderResp)

	c.JSON(http.StatusCreated, orderResp)
}

func GetOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order ID",
		})
		return
	}

	for _, order := range orders {
		if order.ID == id {
			c.JSON(http.StatusOK, order)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Order not found",
	})
}

func UpdateOrderStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order ID",
		})
		return
	}

	var statusReq struct {
		Status models.OrderStatus `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&statusReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid status data",
		})
		return
	}

	for i, order := range orders {
		if order.ID == id {
			orders[i].Status = statusReq.Status
			c.JSON(http.StatusOK, orders[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Order not found",
	})
}