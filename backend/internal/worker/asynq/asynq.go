package asynq

import "github.com/hibiken/asynq"

type IAsynqWorker interface {
	EnqueueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}
