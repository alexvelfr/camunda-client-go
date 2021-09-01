package worker

import "github.com/alexvelfr/camunda-client-go/v2/processor"

type WorkerConfig struct {
	Handlers []*TopicHandler
	Logger   func(err error)
}

type TopicHandler struct {
	Topic   string
	Handler func(ctx *processor.Context) error
}
