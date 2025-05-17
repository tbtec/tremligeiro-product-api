package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type controller struct{}

func NewLivenessController() *controller {
	return &controller{}
}

type Output struct {
	Status string `json:"status"`
}

func (ctrl *controller) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {
	return httpserver.Ok(Output{Status: "OK"})
}
