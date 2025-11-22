package server

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/student-nirajkumar/shopping-challenge/backend/internal/models"
)

func initCartRoutes(rg *gin.RouterGroup) {

    rg.POST("/carts", addItemToCartHandler)
    rg.GET("/carts/me", getMyCartHandler)
    rg.GET("/carts", listAllCartsHandler)
}

// ------------------------------
// POST /carts
// Add item to user's cart
// ------------------------------
type AddItemRequest struct {
    ItemID uint `json:"item_id"`
}

func addItemToCartHandler(c *gin.Context) {
    user := c.MustGet("user").(models.User)

    var body AddItemRequest
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
        return
    }

    // Step 1: If user has no cart, create one
    var cart models.Cart
    if err := db.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
        // Create new cart
        cart = models.Cart{
            UserID: &user.ID,   // ✅ FIXED HERE
            Status: "open",
        }
        db.Create(&cart)
    }

    // Step 2: Add item to cart_items table
    cartItem := models.CartItem{
        CartID: cart.ID,
        ItemID: body.ItemID,
    }
    db.Create(&cartItem)

    c.JSON(http.StatusCreated, gin.H{
        "cart_id": cart.ID,
        "item_id": body.ItemID,
    })
}

// ------------------------------
// GET /carts/me
// ------------------------------
func getMyCartHandler(c *gin.Context) {
    user := c.MustGet("user").(models.User)

    var cart models.Cart
    if err := db.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
        return
    }

    // Load cart items
    var items []models.CartItem
    db.Where("cart_id = ?", cart.ID).Find(&items)

    c.JSON(http.StatusOK, gin.H{
        "cart":  cart,
        "items": items,
    })
}

// ------------------------------
// GET /carts
// ------------------------------
func listAllCartsHandler(c *gin.Context) {
    var carts []models.Cart
    db.Find(&carts)
    c.JSON(http.StatusOK, carts)
}
