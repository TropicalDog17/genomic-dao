package handler

import (
	"net/http"

	"github.com/TropicalDog17/genomic-dao-service/internal/auth"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService *auth.AuthService
}

type AuthHandler interface {
	Register(c *gin.Context)
}

func NewAuthHandler(authService *auth.AuthService) AuthHandler {
	return &authHandler{
		authService,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	// Get the public key from the request body
	var req struct {
		Pubkey []byte `json:"pubkey"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register the user
	userID, err := h.authService.Register(req.Pubkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userID": userID})
}
