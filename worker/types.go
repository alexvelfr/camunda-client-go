package worker

import (
	"time"

	"log"

	"github.com/alexvelfr/camunda-client-go/v2/processor"
)

type WorkerConfig struct {
	Handlers                  []*TopicHandler
	Logger                    func(err error)
	MaxTasks                  int
	MaxParallelTaskPerHandler int
	LockDuration              time.Duration
	LongPollingTimeout        time.Duration
}

func DefaultConfig() WorkerConfig {
	return WorkerConfig{
		Handlers:                  make([]*TopicHandler, 0),
		Logger:                    logErr,
		MaxTasks:                  10,
		MaxParallelTaskPerHandler: 100,
		LockDuration:              time.Second * 5,
		LongPollingTimeout:        time.Second * 5,
	}
}

func logErr(err error) {
	log.Println(err)
}

type TopicHandler struct {
	Topic   string
	Handler func(ctx *processor.Context) error
}
