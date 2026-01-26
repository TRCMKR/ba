package closer

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

type Closer struct {
	ready          chan struct{}
	once           sync.Once
	closeCallbacks []func() error
}

// TODO: add logger
// TODO: add options

// New creates a new Closer instance
// if ctx is canceled, it automatically calls CloseAll
// if signals are passed, it will automatically listen for them and call CloseAll when received
func New(ctx context.Context, signals ...os.Signal) *Closer {
	closer := &Closer{}

	go func() {
		var sigChan chan os.Signal

		if len(signals) > 0 {
			sigChan = make(chan os.Signal, 1)
			signal.Notify(sigChan, signals...)
		}

		select {
		case <-ctx.Done():
			closer.CloseAll()

		// if no signals were not passed, this case will be ignored
		case <-sigChan:
			closer.CloseAll()
		}
	}()

	return closer
}

func (c *Closer) Add(cb func() error) {
	c.closeCallbacks = append(c.closeCallbacks, cb)
}

// Wait waits until every resource is closed
func (c *Closer) Wait() {
	<-c.ready
}

// CloseAll closes all resources
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.ready)
		for _, cb := range c.closeCallbacks {
			cb()
		}
	})
}
