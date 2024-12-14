package server

import (
	"fmt"
	"sync"

	"github.com/TropicalDog17/genomic-dao-service/internal/auth"
	"github.com/TropicalDog17/genomic-dao-service/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var startOnce sync.Once

func StartService(db *gorm.DB) {
	startOnce.Do(func() {
		r := route(db)
		err := r.Run(":8080")
		if err != nil {
			panic(err)
		}
		fmt.Println("Server started at :8080")
	})
}

// Routing by gin
func route(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	authService := auth.NewAuthService(auth.NewUserRepository(db))
	// storageService := storage.NewGenDataRepository(db)

	authHandler := handler.NewAuthHandler(authService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	// Create a new user
	r.POST("/user", authHandler.Register)

	return r
}
