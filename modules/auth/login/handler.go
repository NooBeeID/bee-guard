package login

import (
	"context"
	"net/http"

	"github.com/NooBeeID/bee-guard/infra/contracts"
)

type contractService interface {
	Login(ctx context.Context, req Request) (Response, error)
}
type handler struct {
	svc contractService
}

func NewHandler(svc contractService) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) Handle(ctx context.Context, req contracts.Request) contracts.Response {

	var request Request
	if err := req.ParseRequest(&request); err != nil {
		return contracts.Response{HttpStatus: http.StatusBadRequest, Err: err, Message: err.Error()}
	}

	resp, err := h.svc.Login(ctx, request)
	if err != nil {
		return contracts.Response{Err: err}
	}
	return contracts.Response{
		Payload: resp,
	}
}
