package sdkhttp

import "backend/internal/services"

type IServer struct {
	IdempotencyService services.IIdempotencyService
	AccountsService    services.IAccountsService
	// Add service interfaces here
	// ...
}

func NewServer(
	idempotencyService services.IIdempotencyService,
	accountsService services.IAccountsService,
) *IServer {

	return &IServer{
		IdempotencyService: idempotencyService,
		AccountsService:    accountsService,
	}
}

var Server *IServer
