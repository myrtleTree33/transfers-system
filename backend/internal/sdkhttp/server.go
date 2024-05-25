package sdkhttp

import "backend/internal/services"

type IServer struct {
	IdempotencyService  services.IIdempotencyService
	AccountsService     services.IAccountsService
	TransactionsService services.ITransactionsService
	// Add service interfaces here
	// ...
}

func NewServer(
	idempotencyService services.IIdempotencyService,
	accountsService services.IAccountsService,
	transactionsService services.ITransactionsService,
) *IServer {

	return &IServer{
		IdempotencyService:  idempotencyService,
		AccountsService:     accountsService,
		TransactionsService: transactionsService,
	}
}

var Server *IServer
