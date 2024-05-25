package middleware

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
	"github.com/segmentio/ksuid"
)

// WithAPIKey is a middleware that checks if the API key is valid and if it has access to the resource
func WithIdempotency(
	idempotencyService services.IIdempotencyService,
) gin.HandlerFunc {

	return func(c *gin.Context) {
		keyHash, err := idempotencyService.GenKeyHash(c)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to generate key hash.  Error = " + err.Error()})
			return
		}

		// Retrieve idempotency record from database
		existingIdempotency, err := idempotencyService.GetByKeyHash(keyHash)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to retrieve idempotency record"})
			return
		}

		if existingIdempotency != nil {
			fmt.Println("existingIdempotency", existingIdempotency)
			// TODO get from DB, and return previous response here
			abortAndDecorateWithIdempotentResponse(c, existingIdempotency)
			return
		}

		// Create new blank idempotency record
		err = idempotencyService.CreateOne(&models.Idempotency{
			BaseImmutable: models.BaseImmutable{
				ID: ksuid.New(),
			},
			KeyHash: keyHash,
			HttpResponseHeaders: pgtype.JSONB{
				Bytes:  []byte("{}"),
				Status: pgtype.Present,
			},
		})
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to create idempotency record", "err": err.Error()})
			return
		}

		// save idempotency key hash to context
		c.Set("idempotency_keyhash", keyHash)

		defer func() {
		}()

		c.Next()
	}
}

func abortAndDecorateWithIdempotentResponse(c *gin.Context, idempotency *models.Idempotency) {
	httpResponseCode := idempotency.HttpResponseCode

	if httpResponseCode >= 200 && httpResponseCode < 300 {
		httpResponseCode = http.StatusAlreadyReported
	}

	// Write response
	c.Data(
		httpResponseCode,
		"application/json; charset=utf-8",
		[]byte(idempotency.HttpResponseBody),
	)

	// Abort
	c.Abort()
}
