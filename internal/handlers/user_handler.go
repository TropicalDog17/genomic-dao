package controllers

import (
	db "github.com/TropicalDog17/genomic-dao-service/internal/db"
	"github.com/TropicalDog17/genomic-dao-service/internal/model"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tx := db.Instance.Create(&user)
	if tx.Error != nil {
		context.JSON(400, gin.H{"error": tx.Error.Error()})
		return
	}

}
