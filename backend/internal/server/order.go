package server

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/student-nirajkumar/shopping-challenge/backend/internal/models"
)

func initOrderRoutes(rg *gin.RouterGroup) {
    rg.POST("/orders", createOrderHandler)
    rg.GET("/orders/me", listMyOrdersHandler)
    rg.GET("/orders", listAllOrdersHandler)
}

// ----------------------------------------------
// POST /orders  → Convert cart to an order
// ----------------------------------------------
func createOrderHandler(c *gin.Context) {
    user := c.MustGet("user").(models.User)

    // Find the user's cart
    var cart models.Cart
    if err := db.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "no cart found"})
        return
    }

    // Create new order
    order := models.Order{
        CartID: cart.ID,
        UserID: user.ID,
    }
    db.Create(&order)

    // Optional: mark cart as ordered
    db.Model(&cart).Update("status", "ordered")

    c.JSON(http.StatusCreated, gin.H{
        "order_id": order.ID,
    })
}

// ----------------------------------------------
// GET /orders/me → Logged-in user's orders
// ----------------------------------------------
func listMyOrdersHandler(c *gin.Context) {
    user := c.MustGet("user").(models.User)

    var orders []models.Order
    db.Where("user_id = ?", user.ID).Find(&orders)

    c.JSON(http.StatusOK, orders)
}

// ----------------------------------------------
// GET /orders → All orders (admin)
// ----------------------------------------------
func listAllOrdersHandler(c *gin.Context) {
    var orders []models.Order
    db.Find(&orders)

    c.JSON(http.StatusOK, orders)
}
