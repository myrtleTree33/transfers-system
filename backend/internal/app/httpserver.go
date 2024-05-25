package app

import (
	"context"
	"net/http"
	"time"
)

type httpserver struct {
	server *http.Server
	logger ILogger
}

// NewHttpServer is to initiate http-server instance
func NewHttpServer(addr string, handlers http.Handler, logger ILogger) *httpserver {
	mux := http.NewServeMux()
	mux.Handle("/", handlers)

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &httpserver{
		server: s,
		logger: logger,
	}
}

// Start is to start the http server
func (hs *httpserver) Start() {
	go func() {
		if err := hs.server.ListenAndServe(); err != http.ErrServerClosed {
			hs.logger.FatalfContext(context.Background(), "Failed to run server. err=%v", err)
		}
	}()
}

// Shutdown is to shutdown http server gracefully
func (hs *httpserver) Shutdown() {
	// the reason of having this is
	// to avoid possibly cutting
	// running process

	// giving 10 second before actually shutting down http server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// shutdown http server actual
	hs.server.Shutdown(ctx)
}
