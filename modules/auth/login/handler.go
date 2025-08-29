package login

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NooBeeID/bee-guard/infra/contracts"
)

type ContractService interface {
	Login(ctx context.Context, req Request) (Response, error)
}
type handler struct {
	svc ContractService
}

func NewHandler(svc ContractService) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) Handle(ctx context.Context, req contracts.Request) contracts.Response {
	fmt.Println("Login Handle")
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
