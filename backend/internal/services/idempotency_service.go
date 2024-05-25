package services

import (
	"backend/internal/models"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IIdempotencyService interface {
	CreateOne(idempotency *models.Idempotency) error
	GetByKeyHash(keyHash string) (*models.Idempotency, error)
	UpdateOneByID(idempotecy *models.Idempotency) (*models.Idempotency, error)
	GenKeyHash(c *gin.Context) (string, error)
}

type IdempotencyService struct {
	db *gorm.DB
}

func NewIdempotencyService(db *gorm.DB) IIdempotencyService {
	return &IdempotencyService{
		db: db,
	}
}

func (s *IdempotencyService) CreateOne(idempotency *models.Idempotency) error {
	err := s.db.Create(idempotency).Error
	if err != nil {
		return fmt.Errorf("unable to create idempotency err=%w", err)
	}
	return nil
}

func (s *IdempotencyService) GetByKeyHash(keyHash string) (*models.Idempotency, error) {
	found := &models.Idempotency{}
	err := s.db.Model(models.Idempotency{}).Where("key_hash = ?", keyHash).First(found).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to find idempotency by key hash err=%w", err)
	}
	return found, nil
}

func (s *IdempotencyService) UpdateOneByID(updated *models.Idempotency) (*models.Idempotency, error) {
	err := s.db.Model(models.Idempotency{}).Where("id = ?", updated.ID).Updates(&updated).Error
	if err != nil {
		return updated, fmt.Errorf("unable to update idempotency record err=%w", err)
	}
	return updated, nil
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
