package service

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/uframework/httpserver"
	"github.com/thienhaole92/uframework/notifylog"
)

type HealthCheckRequest struct{}

type HealthHandler struct {
	log notifylog.NotifyLog
}

func (s *Service) CheckHealth(ctx echo.Context, req *HealthCheckRequest) (any, *echo.HTTPError) {
	delegator := func(
		logger notifylog.NotifyLog,
		ctx echo.Context,
		req *HealthCheckRequest,
	) (*httpserver.Response, *echo.HTTPError) {
		handler := NewHealthHandler(logger)

		return handler.Handle(ctx, req)
	}

	return httpserver.Call(ctx, req, "CheckHealth", delegator)
}

func NewHealthHandler(log notifylog.NotifyLog) *HealthHandler {
	return &HealthHandler{
		log: log,
	}
}

func (h *HealthHandler) Handle(_ echo.Context, _ *HealthCheckRequest) (*httpserver.Response, *echo.HTTPError) {
	return &httpserver.Response{
		RequestID: "",
		Data: map[string]any{
			"msg":        "Mon 18 Aug",
			"commit_sha": os.Getenv("COMMIT_SHA"),
		},
		Pagination: nil,
	}, nil
}
