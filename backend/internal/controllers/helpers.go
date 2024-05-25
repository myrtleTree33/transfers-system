package controllers

import (
	"backend/internal/app"
	"backend/internal/models"
	"backend/internal/sdkhttp"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReplyJSONWithIdempotency is a middleware that replies with JSON and updates an idempotency record
func ReplyJSONWithIdempotency(c *gin.Context, httpStatusCode int, httpBody interface{}) {
	err := UpdateIdempotency(c, httpStatusCode, httpBody)
	if err != nil {
		app.Logger.ErrorfContext(c, "unable to update idempotency record err=%w", err)
	}

	c.JSON(httpStatusCode, httpBody)
}

// UpdateIdempotency is a middleware that updates an idempotency record
func UpdateIdempotency(c *gin.Context, httpResponseCode int, httpResponseBody interface{}) error {
	// Retrieve idempotency key from header
	key := c.GetString("idempotency_keyhash")
	if key == "" {
		return fmt.Errorf("idempotency key not found")
	}

	// // Retrieve idempotency from database
	// idempotency, err := sdkhttp.Server.IdempotencyService.FindOneByKeyHashAndOrganisationID(key, org.ID)
	// if err != nil {
	// 	return fmt.Errorf("unable to get idempotency record err=%w", err)
	// }

	// if idempotency == nil {
	// 	return fmt.Errorf("idempotency record not found")
	// }

	// TODO fill this in
	idempotency := &models.Idempotency{}

	// Convert HTTP response body to JSON string
	if httpResponseBody != nil {
		httpBody, err := json.Marshal(httpResponseBody)
		if err != nil {
			return fmt.Errorf("unable to marshal http response body err=%w", err)
		}
		idempotency.HttpResponseBody = string(httpBody)
	}

	// Update idempotency
	idempotency.HttpResponseCode = httpResponseCode
	// idempotency.HttpResponseHeaders = httpResponseHeaders

	_, err := sdkhttp.Server.IdempotencyService.UpdateOneByID(idempotency)
	if err != nil {
		return fmt.Errorf("unable to update idempotency record err=%w", err)
	}

	return nil
}
