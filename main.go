package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	ID          int        `json:"id" gorm:"column:id"`
	Description string     `json:"description" gorm:"column:description"`
	Title       string     `json:"title" gorm:"column:title"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	Updated     *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
type CreateTodoItem struct {
	ID          int    `json:"-" gorm:"column:id;"`
	Description string `json:"description" gorm:"column:description;"`
	Title       string `json:"title" gorm:"column:title;"`
	// Status      string `json:"status" gorm:"column:status;"`
}

func main() {
	godotenv.Load()
	now := time.Now().UTC()
	var fakeData = TodoItem{
		ID:          1,
		Description: "This is a description",
		Title:       "This is a title",
		Status:      "active",
		CreatedAt:   &now,
		Updated:     &now,
	}
	dsn := os.Getenv("DB_SECRET")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database", db)
	// Migrate the schema
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id")
			items.DELETE("/:id")

		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": fakeData,
		})
	})
	r.Run(":9292")
}

func (CreateTodoItem) TableName() string {
	return TodoItem{}.TableName()

}
func (TodoItem) TableName() string {
	return "todo_items"
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var body CreateTodoItem
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&body).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": body.ID})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var body TodoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// body.ID = id
		// db.First(&body, id)
		if err := db.Where("id = ?", id).First(&body).First(&body).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": body})

	}
}
