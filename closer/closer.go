package closer

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"go.uber.org/zap"
)

// TODO: add support for MaxParallel
type config struct {
	MaxParallel uint32
}

type Closer struct {
	ready          chan struct{}
	once           sync.Once
	closeCallbacks []func() error

	cfg *config

	logger *zap.Logger
}

// New creates a new Closer instance
// if ctx is canceled, it automatically calls CloseAll
// if signals are passed, it will automatically listen for them and call CloseAll when received
// by default uses zap.NewNop(), use custom logger with WithOptions
func New(ctx context.Context, signals ...os.Signal) *Closer {
	closer := &Closer{
		ready: make(chan struct{}),
		once:  sync.Once{},
	}

	closer.WithOptions(defaultOptions())

	go func() {
		var sigChan chan os.Signal

		if len(signals) > 0 {
			sigChan = make(chan os.Signal, 1)
			signal.Notify(sigChan, signals...)
		}

		select {
		case <-ctx.Done():
			closer.CloseAll()

		// if no signals were passed, this case will be ignored
		case <-sigChan:
			closer.CloseAll()
		}
	}()

	return closer
}

// Add adds a new close callback
func (c *Closer) Add(cb func() error) {
	c.closeCallbacks = append(c.closeCallbacks, cb)
}

// Wait waits until every resource is closed
func (c *Closer) Wait() {
	<-c.ready
}

// CloseAll closes all resources in reverse order
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.ready)
		c.logger.Info("closing all resources")

		callbacksCount := len(c.closeCallbacks)
		for i, _ := range c.closeCallbacks {
			err := c.closeCallbacks[callbacksCount-i-1]()
			if err != nil {
				c.logger.Error("error during closing",
					zap.Error(err),
				)
			}
		}
	})
}

// WithOptions use custom options to modify closer
func (c *Closer) WithOptions(options ...Option) {
	for _, option := range options {
		option(c)
	}
}
