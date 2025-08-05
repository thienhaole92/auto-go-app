package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/thienhaole92/auto-go-app/internal/config"
	"github.com/thienhaole92/auto-go-app/internal/service"
	"github.com/thienhaole92/uframework/container"
	"github.com/thienhaole92/uframework/httpserver"
	"github.com/thienhaole92/uframework/metricserver"
	"github.com/thienhaole92/uframework/notifylog"
	"github.com/thienhaole92/uframework/runner"
)

const (
	httpSubsystem = "echo"
	metricsPath   = "/metrics"
	statusPath    = "/status"
)

func main() {
	logger := notifylog.New("ws-service", notifylog.JSON)

	cfg, err := config.New(nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load configuration")
	}

	svc := service.New()

	rnn := runner.New(
		runner.WithServer(setupMetricServer(cfg)),
		runner.WithHTTPServer(setupHTTPServer(cfg)),
		runner.WithRestAPIService(setupRestAPIService(svc)),
	)

	rnn.Run()
}

func setupHTTPServer(cfg *config.Config) func(*container.Container) *httpserver.Server {
	return func(_ *container.Container) *httpserver.Server {
		serverOptions := &httpserver.Option{
			Host:             cfg.HTTPServerHost,
			Port:             cfg.HTTPServerPort,
			EnableCors:       cfg.HTTPEnableCORS,
			BodyLimit:        cfg.HTTPBodyLimit,
			ReadTimeout:      cfg.HTTPServerReadTimeout,
			WriteTimeout:     cfg.HTTPServerWriteTimeout,
			GracePeriod:      cfg.GracefulShutdownPeriod,
			Subsystem:        httpSubsystem,
			RequireRequestID: cfg.HTTPSkipRequestID,
		}

		return httpserver.New(serverOptions)
	}
}

func setupMetricServer(cfg *config.Config) *metricserver.Server {
	serverOptions := &metricserver.Option{
		Host:         cfg.MetricServerHost,
		Port:         cfg.MetricServerPort,
		ReadTimeout:  cfg.MetricServerReadTimeout,
		WriteTimeout: cfg.MetricServerWriteTimeout,
		GracePeriod:  cfg.GracefulShutdownPeriod,
		MetricPath:   metricsPath,
		StatusPath:   statusPath,
	}

	return metricserver.New(serverOptions)
}

func setupRestAPIService(svc *service.Service) func(*container.Container) {
	return func(container *container.Container) {
		root := container.MustEchoGroup()

		root.GET("/health", httpserver.Wrapper(svc.CheckHealth))
	}
}
