package server

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/student-nirajkumar/shopping-challenge/backend/internal/models"
)

// POST /items — create item
func createItemHandler(c *gin.Context) {
    var body struct {
        Name   string `json:"name"`
        Status string `json:"status"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
        return
    }

    item := models.Item{
        Name:   body.Name,
        Status: body.Status,
    }

    if err := models.DB.Create(&item).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create item"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "item_id": item.ID,
        "name":    item.Name,
        "status":  item.Status,
    })
}

// GET /items — list items
func listItemsHandler(c *gin.Context) {
    var items []models.Item
    if err := models.DB.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch items"})
        return
    }

    c.JSON(http.StatusOK, items)
}

// Register item routes
func initItemRoutes(auth *gin.RouterGroup) {
    auth.POST("/items", createItemHandler)
    auth.GET("/items", listItemsHandler)
}
