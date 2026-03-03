package models

// OrderStatus represents the status of an order
type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusConfirmed OrderStatus = "CONFIRMED"
	StatusShipped   OrderStatus = "SHIPPED"
	StatusDelivered OrderStatus = "DELIVERED"
	StatusCancelled OrderStatus = "CANCELLED"
)

// OrderRequest represents the structure for creating an order
type OrderRequest struct {
	CustomerID string                 `json:"customerId" binding:"required"`
	Items      []OrderItem            `json:"items" binding:"required"`
	Total      float64                `json:"total" binding:"required"`
	Address    string                 `json:"address" binding:"required"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ItemID   string  `json:"itemId" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
}

// OrderResponse represents the structure for order responses
type OrderResponse struct {
	ID         int                    `json:"id"`
	CustomerID string                 `json:"customerId"`
	Items      []OrderItem            `json:"items"`
	Total      float64                `json:"total"`
	Address    string                 `json:"address"`
	Status     OrderStatus            `json:"status"`
	CreatedAt  string                 `json:"createdAt,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}
