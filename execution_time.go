package executiontime

import (
	"time"

	"github.com/gin-gonic/gin"
)

const defaultHeaderKey = "X-Execution-Time"

// Option for execution time
type Option func(*config)

func New(opts ...Option) gin.HandlerFunc {
	cfg := &config{
		headerKey: defaultHeaderKey,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(g *gin.Context) {
		begin := time.Now()
		g.Next()
		cost := time.Since(begin)

		g.Header(cfg.headerKey, cost.String())
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
