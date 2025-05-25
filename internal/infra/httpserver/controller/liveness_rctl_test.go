package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

func TestNewLivenessController(t *testing.T) {
	ctrl := NewLivenessController()
	assert.NotNil(t, ctrl)
}

func TestLivenessController_Handle(t *testing.T) {
	ctrl := NewLivenessController()
	req := httpserver.Request{} // Assuming no fields required for liveness
	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 200, resp.Code)
	output, ok := resp.Body.(Output)
	assert.True(t, ok)
	assert.Equal(t, "OK", output.Status)
}
