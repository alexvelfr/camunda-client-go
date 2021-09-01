package worker

import (
	"context"
	"time"

	camunda_client_go "github.com/alexvelfr/camunda-client-go/v2"
	"github.com/alexvelfr/camunda-client-go/v2/processor"
)

// Wrapper for fast create camunda workers
type Worker struct {
	handlers []*TopicHandler
	client   *camunda_client_go.Client
	logger   func(err error)
}

// Run run worker. Run is blocking func
func (a *Worker) Run(ctx context.Context) {
	proc := processor.NewProcessor(a.client, &processor.ProcessorOptions{
		WorkerId:                  "worker",
		LockDuration:              time.Second * 5,
		MaxTasks:                  10,
		MaxParallelTaskPerHandler: 100,
		LongPollingTimeout:        5 * time.Second,
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
func NewWorker(client *camunda_client_go.Client, config WorkerConfig) *Worker {
	return &Worker{
		handlers: config.Handlers,
		client:   client,
		logger:   config.Logger,
	}
}
