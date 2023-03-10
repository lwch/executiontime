package executiontime

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultHeaderKey = "X-Execution-Time"

// Option for execution time
type Option func(*config)

type writer struct {
	gin.ResponseWriter
	key   string
	begin time.Time
	once  sync.Once
}

func (w *writer) Write(p []byte) (int, error) {
	w.once.Do(func() {
		w.Header().Set(w.key, time.Since(w.begin).String())
	})
	return w.ResponseWriter.Write(p)
}

func New(opts ...Option) gin.HandlerFunc {
	cfg := &config{
		headerKey: defaultHeaderKey,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(g *gin.Context) {
		g.Writer = &writer{
			ResponseWriter: g.Writer,
			key:            cfg.headerKey,
			begin:          time.Now(),
		}
		g.Next()
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
