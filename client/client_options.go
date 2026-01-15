package client

import "time"

type Option func(*config)

// WithTimeout sets the timeout duration for client requests.
func WithTimeout(timeout int64, size time.Duration) Option {
	return func(c *config) {
		c.timeout = size * time.Duration(timeout)
	}
}

// WithDefaults sets default values for the client configuration.
func WithDefaults() Option {
	return func(c *config) {
		if c.timeout == 0 {
			c.timeout = 5 * time.Second
		}
	}
}
