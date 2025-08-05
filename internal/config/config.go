package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// Common settings
	GracefulShutdownPeriod time.Duration `env:"GRACEFUL_SHUTDOWN_PERIOD" envDefault:"30s"`

	// Metric server settings
	MetricServerHost         string        `env:"METRIC_SERVER_HOST"          envDefault:"0.0.0.0"`
	MetricServerPort         int           `env:"METRIC_SERVER_PORT"          envDefault:"8082"`
	MetricServerReadTimeout  time.Duration `env:"METRIC_SERVER_READ_TIMEOUT"  envDefault:"30s"`
	MetricServerWriteTimeout time.Duration `env:"METRIC_SERVER_WRITE_TIMEOUT" envDefault:"30s"`

	// HTTP server settings
	HTTPServerHost         string        `env:"HTTP_SERVER_HOST"          envDefault:"0.0.0.0"`
	HTTPServerPort         int           `env:"HTTP_SERVER_PORT"          envDefault:"8081"`
	HTTPServerReadTimeout  time.Duration `env:"HTTP_SERVER_READ_TIMEOUT"  envDefault:"30s"`
	HTTPServerWriteTimeout time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT" envDefault:"30s"`
	HTTPEnableCORS         bool          `env:"HTTP_ENABLE_CORS"          envDefault:"false"`
	HTTPBodyLimit          string        `env:"HTTP_BODY_LIMIT"           envDefault:"100K"`
	HTTPSkipRequestID      bool          `env:"HTTP_SKIP_REQUEST_ID"      envDefault:"true"`
}

func New(opts *env.Options) (*Config, error) {
	cfg := new(Config)
	if opts != nil {
		if err := env.ParseWithOptions(cfg, *opts); err != nil {
			return nil, err
		}

		return cfg, nil
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
