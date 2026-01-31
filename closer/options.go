package closer

import (
	"go.uber.org/zap"
)

// Option is type for changing default closer options
type Option func(*Closer)

// WithLogger set custom logger
func WithLogger(logger *zap.Logger) Option {
	childLogger := logger.With(
		zap.String("component", "closer"),
	)

	return func(c *Closer) {
		c.logger = childLogger
	}
}

// WithMaxParallel set max amount of parallel closings
func WithMaxParallel(parallel uint32) Option {
	return func(c *Closer) {
		c.cfg.MaxParallel = parallel
	}
}

// defaultOptions default closer options
func defaultOptions() Option {
	return func(c *Closer) {
		c.cfg = &config{
			MaxParallel: 1,
		}
		c.logger = zap.NewNop()
	}
}
