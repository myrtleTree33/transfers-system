package middleware

import (
	"backend/internal/models"
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
)

// WithAPIKey is a middleware that checks if the API key is valid and if it has access to the resource
func WithIdempotency(
	idempotencyService services.IIdempotencyService,
) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Retrieve secret key from header
		secretKey := c.GetHeader("Authorization")
		if secretKey == "" {
			c.AbortWithStatusJSON(500, gin.H{"msg": "API secret key not found"})
			return
		}

		keyHash, err := idempotencyService.GenKeyHash(c)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to generate key hash.  Error = " + err.Error()})
			return
		}

		// // Retrieve idempotency record from database
		// existingIdempotency, err := idempotencyService.FindOneByKeyHashAndOrganisationID(keyHash, org.ID)
		// if err != nil {
		// 	c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to retrieve idempotency record"})
		// 	return
		// }

		// if existingIdempotency != nil {
		// 	fmt.Println("existingIdempotency", existingIdempotency)
		// 	// TODO get from DB, and return previous response here
		// 	abortAndDecorateWithIdempotentResponse(c, existingIdempotency)
		// }

		// Create new blank idempotency record
		idempotencyService.CreateOne(&models.Idempotency{
			KeyHash: keyHash,
			HttpResponseHeaders: pgtype.JSONB{
				Bytes:  []byte("{}"),
				Status: pgtype.Present,
			},
		})

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
