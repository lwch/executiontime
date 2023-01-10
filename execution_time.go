package executiontime

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

var headerXExecutionTime string

// Option for execution time
type Option func(*config)

func New(opts ...Option) app.HandlerFunc {
	cfg := &config{
		headerKey: "X-Execution-Time",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(ctx context.Context, c *app.RequestContext) {
		begin := time.Now()
		c.Next(ctx)
		cost := time.Since(begin)

		headerXExecutionTime = cfg.headerKey
		c.Header(headerXExecutionTime, cost.String())
	}
}

// WithCustomHeaderStrKey set custom header key for execution time
func WithCustomHeaderStrKey(key string) Option {
	return func(cfg *config) {
		cfg.headerKey = key
	}
}

// Config defines the config for ExecutionTime middleware
type config struct {
	headerKey string
}
