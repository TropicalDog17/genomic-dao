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

// @BasePath /

// Register godoc
// @Summary Register a new user
// @Description Register a new user with their public key
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "User registration details"
// @Success 200 {object} RegisterResponse "Successfully registered"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	// Get the public key from the request body
	var req struct {
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register the user
	userID, err := h.authService.Register(req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userID": userID})
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Address string `json:"address" example:"0x6491414173c71986Ee031307Af447cE1DbDf2ED0"` // Public key in bytes
}

// RegisterResponse represents the response for successful registration
type RegisterResponse struct {
	UserID string `json:"userID" example:"user123"` // Unique identifier for the registered user
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"invalid request format"` // Error message
}
