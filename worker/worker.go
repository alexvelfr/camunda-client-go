package worker

import (
	"context"
	"time"

	camunda_client_go "github.com/alexvelfr/camunda-client-go/v2"
	"github.com/alexvelfr/camunda-client-go/v2/processor"
)

// Wrapper for fast create camunda workers
type Worker struct {
	handlers                  []*TopicHandler
	client                    *camunda_client_go.Client
	logger                    func(err error)
	maxTasks                  int
	maxParallelTaskPerHandler int
	lockDuration              time.Duration
	longPollingTimeout        time.Duration
	name                      string
}

// Run run worker. Run is blocking func
func (a *Worker) Run(ctx context.Context) {
	proc := processor.NewProcessor(a.client, &processor.ProcessorOptions{
		WorkerId:                  a.name,
		LockDuration:              a.lockDuration,
		MaxTasks:                  a.maxTasks,
		MaxParallelTaskPerHandler: a.maxParallelTaskPerHandler,
		LongPollingTimeout:        a.longPollingTimeout,
	}, a.logger)
	for _, h := range a.handlers {
		proc.AddHandler(
			&[]camunda_client_go.QueryFetchAndLockTopic{
				{TopicName: h.Topic},
			},
			h.Handler,
		)
	}
	<-ctx.Done()
}

// NewWorker create new Worker
func NewWorker(client *camunda_client_go.Client, name string, config WorkerConfig) *Worker {
	return &Worker{
		name:                      name,
		handlers:                  config.Handlers,
		client:                    client,
		logger:                    config.Logger,
		maxTasks:                  config.MaxTasks,
		maxParallelTaskPerHandler: config.MaxParallelTaskPerHandler,
		lockDuration:              config.LockDuration,
		longPollingTimeout:        config.LongPollingTimeout,
	}
}
