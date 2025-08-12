package service_test

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/thienhaole92/auto-go-app/internal/service"
	"github.com/thienhaole92/uframework/notifylog"
)

func TestNewHealthHandler(t *testing.T) {
	t.Parallel()

	mockLog := notifylog.New("test-logger", notifylog.JSON)
	handler := service.NewHealthHandler(mockLog)

	require.NotNil(t, handler)
}

func TestHealthHandler_Handle(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mockLog := notifylog.New("test-logger", notifylog.JSON)
			h := service.NewHealthHandler(mockLog)

			got, err := h.Handle(nil, nil)

			require.Equal(t, test.wantData, got.Data)
			require.Equal(t, test.wantErr, err)
		})
	}
}
