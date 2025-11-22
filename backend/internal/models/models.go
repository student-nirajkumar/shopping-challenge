 package models

import (
    "time"

    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(d *gorm.DB) {
    DB = d
}

// ---------------------
// Database Models
// ---------------------

// User model
type User struct {
    ID           uint      `gorm:"primaryKey"`
    Username     string    `gorm:"uniqueIndex;not null"`
    PasswordHash string    `gorm:"not null"`
    Token        string    `gorm:"size:255"`
    CartID       *uint
    CreatedAt    time.Time
}

// Item model
type Item struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Status    string    `gorm:"default:'available'"`
    CreatedAt time.Time
}

// Cart model
type Cart struct {
    ID        uint       `gorm:"primaryKey"`
    UserID    *uint
    Name      string
    Status    string     `gorm:"default:'active'"`
    Items     []CartItem `gorm:"foreignKey:CartID"`
    CreatedAt time.Time
}

// CartItem (join table)
type CartItem struct {
    CartID uint `gorm:"primaryKey"`
    ItemID uint `gorm:"primaryKey"`
}

// Order model
type Order struct {
    ID        uint      `gorm:"primaryKey"`
    CartID    uint
    UserID    uint
    CreatedAt time.Time
}

// ---------------------
// Password helpers
// ---------------------

func HashPassword(password string) (string, error) {
    b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(b), err
}

func CheckPassword(hashed string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
