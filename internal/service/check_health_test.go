package service_test

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/thienhaole92/uframework/notifylog"

	"github.com/thienhaole92/auto-go-app/internal/service"
)

func TestNewHealthHandler(t *testing.T) {
	mockLog := notifylog.New("test-logger", notifylog.JSON)
	handler := service.NewHealthHandler(mockLog)

	require.NotNil(t, handler)
}

func TestHealthHandler_Handle(t *testing.T) {
	tests := []struct {
		name     string
		wantData map[string]any
		wantErr  *echo.HTTPError
	}{
		{
			name:     "successful health check",
			wantData: map[string]any{"msg": "OK"},
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLog := notifylog.New("test-logger", notifylog.JSON)
			h := service.NewHealthHandler(mockLog)

			got, err := h.Handle(nil, nil)

			require.Equal(t, tt.wantData, got.Data)
			require.Equal(t, tt.wantErr, err)
		})
	}
}
