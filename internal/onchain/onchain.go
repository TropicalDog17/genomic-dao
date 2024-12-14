package onchain

import (
	"context"
	"fmt"
	"math/big"

	contracts "github.com/TropicalDog17/genomic-dao-service/internal/onchain/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type OnchainService struct {
	client     *ethclient.Client
	auth       *bind.TransactOpts
	controller *contracts.Controller
	geneNFT    *contracts.GeneNFT
	pcspToken  *contracts.PCSPToken
}

func NewOnchainService(client *ethclient.Client, auth *bind.TransactOpts, controllerAddr common.Address) *OnchainService {
	controller, err := contracts.NewController(controllerAddr, client)
	if err != nil {
		panic(err)
	}

	nftAddr, err := controller.GeneNFT(nil)
	if err != nil {
		panic(err)
	}

	geneNFT, err := contracts.NewGeneNFT(nftAddr, client)
	if err != nil {
		panic(err)
	}

	pcspTokenAddr, err := controller.PcspToken(nil)
	if err != nil {
		panic(err)
	}

	pcspToken, err := contracts.NewPCSPToken(pcspTokenAddr, client)
	if err != nil {
		panic(err)
	}

	return &OnchainService{
		client:     client,
		auth:       auth,
		controller: controller,
		geneNFT:    geneNFT,
		pcspToken:  pcspToken,
	}
}

// UploadData starts the upload session on blockchain
func (s *OnchainService) UploadData(docID string) (string, error) {
	// Call uploadData on controller contract
	tx, err := s.controller.UploadData(s.auth, docID)
	if err != nil {
		return "", fmt.Errorf("failed to initiate upload: %v", err)
	}

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), s.client, tx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction: %v", err)
	}

	// Get session ID from event
	for _, log := range receipt.Logs {
		event, err := s.controller.ParseUploadData(*log)
		if err == nil && event != nil {
			return event.SessionId.String(), nil
		}
	}

	return "", fmt.Errorf("failed to get session ID from event")
}

func (s *OnchainService) ConfirmUpload(docID string, contentHash string, proof string, sessionID string, riskScore int) error {

	// parse sessionID to big.Int
	bigSessionID, ok := new(big.Int).SetString(sessionID, 10)
	bigRiskScore := big.NewInt(int64(riskScore))
	if !ok {
		return fmt.Errorf("failed to parse session ID")
	}
	// Call confirmUpload on controller contract
	tx, err := s.controller.Confirm(s.auth, docID, contentHash, proof, bigSessionID, bigRiskScore)
	if err != nil {
		return fmt.Errorf("failed to confirm upload: %v", err)
	}

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), s.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}

	// Log the transaction hash
	fmt.Printf("Confirmed upload with tx hash: %s\n", tx.Hash().Hex())

	// Process events
	for _, log := range receipt.Logs {
		event, err := s.controller.ParseGeneNFTMinted(*log)
		if err == nil && event != nil {
			fmt.Printf("Minted GeneNFT with token ID: %s\n", event.TokenId.String())
		}

		event2, err := s.controller.ParsePCSPRewarded(*log)
		if err == nil && event2 != nil {
			fmt.Printf("Rewarded PCSP with amount: %s\n", event2.Amount.String())
		}
	}
	return nil
}

func (s *OnchainService) GetSession(sessionID string) (*contracts.ControllerUploadSession, error) {
	// parse sessionID to big.Int
	bigSessionID, ok := new(big.Int).SetString(sessionID, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse session ID")
	}

	// Call getSession on controller contract
	data, err := s.controller.GetSession(nil, bigSessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session data: %v", err)
	}

	return &data, nil
}

func (s *OnchainService) GetPcspTokenBalance(address *common.Address) (*big.Int, error) {
	return s.pcspToken.BalanceOf(nil, *address)
}
