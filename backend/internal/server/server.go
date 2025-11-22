package server

import (
    "net/http"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/gin-gonic/gin"
    "github.com/student-nirajkumar/shopping-challenge/backend/internal/models"
)

var db *gorm.DB

// Start initializes DB, runs migrations, and starts the server.
func Start() error {

    // PostgreSQL connection string
    dsn := "host=localhost user=postgres password=root dbname=shopping port=5432 sslmode=disable"

    // Connect to PostgreSQL
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    // Run migrations
    if err := autoMigrate(db); err != nil {
        return err
    }

    // Inject DB into models package
    models.SetDB(db)

    // Gin router
    r := gin.Default()

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "time":   time.Now(),
        })
    })

    // Public routes (signup, login)
    initAuthRoutes(r)

    // =====================================
    // Protected routes (require auth token)
    // =====================================
    auth := r.Group("/", AuthMiddleware())

    // Cart routes
    initCartRoutes(auth)

    // Item routes  <-- IMPORTANT LINE ADDED HERE
    initItemRoutes(auth)
    initOrderRoutes(auth)

    // =====================================

    // Run server on port 8080
    return r.Run(":8080")
}

func autoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.User{},
        &models.Item{},
        &models.Cart{},
        &models.CartItem{},
        &models.Order{},
    )
}

