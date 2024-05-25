package app

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
)

var Worker *worker

type worker struct {
	redisUri  string
	logger    ILogger
	client    *asynq.Client
	server    *asynq.Server
	mux       *asynq.ServeMux
	scheduler *asynq.Scheduler
	queues    map[string]int
}

// NewWorker is to initiate worker
func NewWorker(logger ILogger, cfg *AppConfig) *worker {
	w := &worker{
		redisUri: cfg.RedisURI,
		logger:   logger,
		mux:      asynq.NewServeMux(),
		queues:   map[string]int{},
	}

	// exec post setup callback
	w.setup()

	return w
}

// constructRedisConnOpt is to construct redis connection option
func (w *worker) constructRedisConnOpt() asynq.RedisClientOpt {
	opt, _ := asynq.ParseRedisURI(w.redisUri)
	clientOpt := opt.(asynq.RedisClientOpt)

	// TODO: replace conn assigning with proper values + other missing values
	return asynq.RedisClientOpt{
		Network:  "tcp",
		Addr:     clientOpt.Addr,
		Username: clientOpt.Username,
		Password: clientOpt.Password,
		DB:       clientOpt.DB,

		// add additional configs
		// ...
	}
}

// setup is to initiate worker's client and scheduler
func (w *worker) setup() {
	// construct redis connection
	conn := w.constructRedisConnOpt()

	// initiate client
	w.client = asynq.NewClient(conn)

	// initiate scheduler
	w.scheduler = asynq.NewScheduler(
		conn,
		&asynq.SchedulerOpts{
			Logger: w.logger,

			// add additional configs
			// ...
		},
	)
}

// Start is to start worker
func (w *worker) Start() {
	w.logger.Info("starting worker...")

	// construct redis connection
	conn := w.constructRedisConnOpt()

	// initiate server
	// moving this to .Start() so queues,
	// can be included on the server as queues were
	// passed post worker's initiatlization
	w.server = asynq.NewServer(
		conn,
		asynq.Config{
			Logger:          w.logger,
			ShutdownTimeout: 1 * time.Minute,
			Queues:          w.queues,

			// add additional configs
			// ...
		},
	)

	// start worker's scheduler
	if err := w.scheduler.Start(); err != nil {
		w.logger.FatalfContext(context.Background(), "found error on starting worker's scheduler. err=%v", err)
	}

	// start worker's server
	if err := w.server.Start(w.mux); err != nil {
		w.logger.FatalfContext(context.Background(), "found error on starting worker's server. err=%v", err)
	}

	w.logger.Info("worker started")
}

// Shutdown is to shutdown worker collectively
func (w *worker) Shutdown() {
	w.scheduler.Shutdown()
	w.server.Shutdown()
	if err := w.client.Close(); err != nil {
		w.logger.WarnfContext(context.Background(), "found error on closing worker's client. err=%v", err)
	}
}

// handler is the expected spec of worker-handler's functionality
type handler interface {
	Handle(context.Context, *asynq.Task) error
	GetName() string
}

// RegisterHandlers is to register worker's handlers
func (w *worker) RegisterHandlers(handlers ...handler) {
	for _, handler := range handlers {
		handler := handler // doing this because of pointer issue in golang
		w.mux.HandleFunc(handler.GetName(), handler.Handle)
	}
}

// scheduler is the expected spec of worker-scheduler's functionality
type scheduler interface {
	GetCronSpec() string
	NewSchedulerTask() *asynq.Task
}

// RegisterScheduler is to register worker's schedulers. Referring to scheduled task
func (w *worker) RegisterSchedulers(schedulers ...scheduler) {
	for _, scheduler := range schedulers {
		scheduler := scheduler // doing this because of pointer issue in golang
		task := scheduler.NewSchedulerTask()

		// attempt to register scheduler
		entryId, err := w.scheduler.Register(scheduler.GetCronSpec(), task)

		// log out the output
		if err != nil {
			w.logger.FatalfContext(context.Background(), "found error on registering '%s' scheduler. err=%v", task.Type(), err)
		} else {
			w.logger.InfofContext(context.Background(), "registered '%s' scheduler at '%s'", task.Type(), entryId)
		}
	}
}

// RegisterQueues is register queue's prioritization. The map's as param will consists
// key containg queue's name in string, and value containing queue's priority in int.
func (w *worker) RegisterQueues(queuePriority map[string]int) {
	w.queues = queuePriority
}

// EnqueueTask is to enqueue task to be processed by worker.
func (w *worker) EnqueueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	if len(opts) > 0 {
		return w.client.Enqueue(task, opts...)
	}

	return w.client.Enqueue(task)
}
