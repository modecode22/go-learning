package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Database model structures
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // Store hashed password in real-world applications
	Type     string `json:"type"`     // "customer", "restaurant", "delivery"
}

type Restaurant struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FoodItem struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Restaurant  Restaurant
	RestaurantID uint
}

type Order struct {
	ID       uint `json:"id"`
	User     User
	UserID   uint
	FoodItem FoodItem
	FoodID   uint
	Status   string `json:"status"` // "preparing", "on the way", "delivered"
}

var db *gorm.DB // Assuming you've set this up elsewhere

func RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(200, user)
}

func ListRestaurants(c *gin.Context) {
	var restaurants []Restaurant
	db.Find(&restaurants)
	c.JSON(200, restaurants)
}

func PlaceOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&order)
	c.JSON(200, order)
}

func main() {
	r := gin.Default()

	r.POST("/register", RegisterUser)
	r.GET("/restaurants", ListRestaurants)
	r.POST("/order", PlaceOrder)

	r.Run(":8080")
}
