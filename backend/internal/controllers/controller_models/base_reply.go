package controller_models

import "backend/internal/models"

type BaseReply struct {
	FailureCode models.FailureCodeType `json:"failure_code,omitempty"`
	Error       string                 `json:"error,omitempty"`
}
