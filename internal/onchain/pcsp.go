package onchain

import (
	"math/big"

	contracts "github.com/TropicalDog17/genomic-dao-service/internal/onchain/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type pcspService struct {
	client    *ethclient.Client
	auth      *bind.TransactOpts
	pcspToken *contracts.PCSPToken
}

type PcspService interface {
	GetTokenBalance(address *common.Address) (*big.Int, error)
}

func NewPcspService(client *ethclient.Client, auth *bind.TransactOpts, pcspTokenAddr common.Address) *pcspService {
	pcspToken, err := contracts.NewPCSPToken(pcspTokenAddr, client)
	if err != nil {
		panic(err)
	}

	return &pcspService{
		client:    client,
		auth:      auth,
		pcspToken: pcspToken,
	}
}

func (s *pcspService) GetTokenBalance(address *common.Address) (*big.Int, error) {
	return s.pcspToken.BalanceOf(nil, *address)
}
