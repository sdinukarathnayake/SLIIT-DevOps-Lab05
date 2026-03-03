package main

import (
	"order-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS middleware for cross-origin requests
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Order routes
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", controllers.GetOrders)
		orderGroup.POST("/", controllers.PlaceOrder)
		orderGroup.GET("/:id", controllers.GetOrder)
		orderGroup.PUT("/:id/status", controllers.UpdateOrderStatus)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "order-service",
		})
	})

	router.Run(":8082")
}
