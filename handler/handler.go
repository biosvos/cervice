package handler

import (
	greeter "cervice/proto"
	"context"
	"fmt"
)

var _ greeter.GreeterHandler = &Handler{}

type Handler struct{}

func (h *Handler) Hello(_ context.Context, request *greeter.Request, response *greeter.Response) error {
	response.Msg = fmt.Sprintf("hello %v", request.GetName())
	return nil
}
