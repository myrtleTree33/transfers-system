package services

import (
	"backend/internal/models"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type IIdempotencyService interface {
	GenKeyHash(c *gin.Context) (string, error)
	FindOneByKeyHashAndOrganisationID(keyHash string, organisationID ksuid.KSUID) (*models.Idempotency, error)
	CreateOne(idempotency *models.Idempotency) error
	UpdateOneByID(idempotecy *models.Idempotency) (*models.Idempotency, error)
}

type IdempotencyService struct {
}

func NewIdempotencyService() IIdempotencyService {
	return &IdempotencyService{}
}

func (s *IdempotencyService) CreateOne(idempotency *models.Idempotency) error {
	// TODO fill in
	return nil
}

func (s *IdempotencyService) FindOneByKeyHashAndOrganisationID(keyHash string, organisationID ksuid.KSUID) (*models.Idempotency, error) {
	// TODO fill in
	return nil, nil
}

func (s *IdempotencyService) UpdateOneByID(idempotecy *models.Idempotency) (*models.Idempotency, error) {
	// TODO fill in
	return nil, nil
}

func (s *IdempotencyService) GenKeyHash(c *gin.Context) (string, error) {
	// Deep copy request body
	requestBodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read request body err=%w", err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
	requestBody := string(requestBodyBytes)

	// Retrieve idempotency key from header
	idempotencyKey := c.GetHeader("idempotency-key")
	if idempotencyKey == "" {
		return "", fmt.Errorf("no idempotency key found in header.  attach an idempotency key to the HTTP request header with key 'idempotency-key'")
	}

	// Retrieve URL endpoint
	url := c.Request.URL.String()

	// Retrieve HTTP method
	method := c.Request.Method

	// Generate key hash
	key := fmt.Sprintf("%s-%s-%s-%s", method, url, requestBody, idempotencyKey)

	// Generate keyhash as SHA-256
	keyHash := sha256.Sum256([]byte(key))
	keyHashString := fmt.Sprintf("%x", keyHash)

	return keyHashString, nil
}
