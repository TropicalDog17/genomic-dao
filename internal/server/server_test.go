package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/TropicalDog17/genomic-dao-service/internal/auth"
	"github.com/TropicalDog17/genomic-dao-service/internal/server"
	"github.com/TropicalDog17/genomic-dao-service/internal/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IntegrationTestSuite struct {
	suite.Suite
	db     *gorm.DB
	client *ethclient.Client
	server *httptest.Server
}

var (
	testPrivKey = "a0eb1aed34398a6ce10a82dd495bc9f82e6b57b6514d46705a2d105551b78289"
	testPubKey  = "0x62f563A2e09c7987dECBFF61fdcC89cd74717721" // derived from the above private key
)

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	// Setup test database
	// load env
	godotenv.Load("../../.env")

	// TODO: pre-fund the test private key
	os.Setenv("PRIVATE_KEY", testPrivKey)

	var err error
	s.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		s.T().Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	err = s.db.AutoMigrate(&auth.User{})
	if err != nil {
		s.T().Fatal("Failed to run migrations:", err)
	}
	err = s.db.AutoMigrate(&storage.GeneData{})
	if err != nil {
		s.T().Fatal("Failed to run migrations:", err)
	}
	// Connect to Ethereum client
	s.client, err = ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		s.T().Fatal("Failed to connect to Ethereum client:", err)
	}

	// Start the server
	go server.StartService(s.db)
	// Wait for server to start
	time.Sleep(2 * time.Second)
}

func (s *IntegrationTestSuite) TestCompleteUserFlow() {
	t := s.T()

	// 1. Test user registration
	registerPayload := map[string]interface{}{
		"address": testPubKey,
	}
	registerBody, _ := json.Marshal(registerPayload)
	resp, err := http.Post("http://localhost:8080/auth/register", "application/json", bytes.NewBuffer(registerBody))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 2. Test genomic data upload using multipart/form-data
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Add genomic data field
	genomicField, err := writer.CreateFormField("genomicData")
	assert.NoError(t, err)
	_, err = genomicField.Write([]byte("test data"))
	assert.NoError(t, err)

	// Add pubkey field
	pubkeyField, err := writer.CreateFormField("pubkey")
	assert.NoError(t, err)
	_, err = pubkeyField.Write([]byte(testPubKey))
	assert.NoError(t, err)

	// Close the writer
	err = writer.Close()
	assert.NoError(t, err)

	// Create and send the request
	req, err := http.NewRequest("POST", "http://localhost:8080/upload", &b)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read and debug print the raw response body
	rawBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()
	fmt.Printf("Raw response body: %s\n", string(rawBody))

	// Create a new reader with the raw body for JSON decoding
	var uploadResponse UploadResponse
	err = json.NewDecoder(bytes.NewReader(rawBody)).Decode(&uploadResponse)
	assert.NoError(t, err)

	// Validate response fields
	assert.NotEmpty(t, uploadResponse.FileID, "FileID should not be empty")
	assert.NotEmpty(t, uploadResponse.SessionID, "SessionID should not be empty")

	// 3. Test genomic data retrieval
	resp, err = http.Get(fmt.Sprintf("http://localhost:8080/retrieve?fileID=%s", uploadResponse.FileID))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 4. Test PCSP balance
	resp, err = http.Get("http://localhost:8080/pcsp/balance?address=" + testPubKey)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var balanceResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&balanceResponse)
	assert.NoError(t, err)
	assert.Contains(t, balanceResponse, "balance")
	fmt.Println("PCSP balance:", balanceResponse["balance"])
}

// UploadResponse struct to match the handler's response structure
type UploadResponse struct {
	SessionID string `json:"sessionID"`
	Message   string `json:"message"`
	FileID    string `json:"fileID"`
}
