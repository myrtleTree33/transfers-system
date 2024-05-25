package main

import (
	"backend/internal/app"
	"backend/internal/controllers"
	"backend/internal/controllers/accounts_routes"
	"backend/internal/middleware"
	"backend/internal/sdkhttp"
	"backend/internal/services"
	"backend/internal/utilities"
	"context"
	"embed"

	_ "backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	//go:embed configs
	embedFS embed.FS
)

// @title          	Transfers System
// @version         1.0
// @description     This is an http-request collection of Transfers System.
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	// setup logger
	app.Logger = app.NewLogger()
	app.Logger.InfofContext(context.Background(), "Starting app...")

	// setup configs
	app.Config = app.NewConfig(embedFS)

	// initiate worker
	app.Worker = app.NewWorker(app.Logger, app.Config)

	// Initialise DB
	db := app.NewDB(app.Config, app.Logger)

	// Initialise services
	idempotencyService := services.NewIdempotencyService()
	accountsService := services.NewAccountsService(db)

	sdkhttp.Server = sdkhttp.NewServer(
		idempotencyService,
		accountsService,
	)

	// setup individual workers
	//...

	// register worker's handlers
	// app.Worker.RegisterHandlers()

	// register worker's schedulers
	// app.Worker.RegisterSchedulers()

	// register worker-queue's prioritization
	// app.Worker.RegisterQueues(map[string]int{})

	// setup http server
	httpserver := app.NewHttpServer(
		app.Config.AppAddress,
		initHandler(sdkhttp.Server),
		app.Logger,
	)

	// start httpserver
	httpserver.Start()

	// run worker
	app.Worker.Start()

	// wait for termination
	utilities.WatchForExitSignal()

	// shutdown http server
	httpserver.Shutdown()

	// shutdown handler
	app.Worker.Shutdown()
}

// initHandler is to setup app's handler
func initHandler(server *sdkhttp.IServer) *gin.Engine {
	// prep middlewares
	// withIdempotency := middleware.WithIdempotency(server.IdempotencyService)

	// setup gin handler
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Use(middleware.WithTraceId())
	{
		// Heartbeat
		v1.GET("/", controllers.Pong)

		// Partner routes
		accounts := v1.Group("/accounts")
		{

			accounts.POST("/",
				accounts_routes.CreateAccountByID,
			)
			accounts.GET("/:account_id",
				accounts_routes.GetAccountByID,
			)
		}
	}

	// spawn swagger ui
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
