package sdkhttp

import "backend/internal/services"

type IServer struct {
	IdempotencyService services.IIdempotencyService
	// Add service interfaces here
	// ...
}

func NewServer(
	idempotencyService services.IIdempotencyService,
) *IServer {

	return &IServer{
		IdempotencyService: idempotencyService,
	}
}

var Server *IServer
