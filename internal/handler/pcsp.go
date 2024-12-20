package handler

import (
	"github.com/TropicalDog17/genomic-dao-service/internal/onchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type PCSPHandler interface {
	GetPCSPBalance(c *gin.Context)
}

type pcspHandler struct {
	pcspService onchain.PcspService
}

func NewPCSPHandler(pcspService onchain.PcspService) PCSPHandler {
	return &pcspHandler{
		pcspService: pcspService,
	}
}

type PCSPBalanceResponse struct {
	Balance string `json:"balance"`
}

// @Summary Get PCSP balance
// @Description Get the balance of PCSP tokens for a user
// @Tags pcsp
// @Accept json
// @Produce json
// @Param address query string true "User's address"
// @Success 200 {object} PCSPBalanceResponse
// @Failure 500 {object} ErrorResponse
// @Router /pcsp/balance [get]
func (h *pcspHandler) GetPCSPBalance(c *gin.Context) {
	address := c.Query("address")
	ethAddress := common.HexToAddress(address)
	balance, err := h.pcspService.GetTokenBalance(&ethAddress)
	if err != nil {
		c.JSON(500, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, gin.H{"balance": balance.String()})
}
