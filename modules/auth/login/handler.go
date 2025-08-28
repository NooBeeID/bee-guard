package login

import (
	"context"
	"fmt"
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

func (h handler) Login(ctx context.Context, req contracts.Request) contracts.Response {
	fmt.Println("Login Handler")

	var request Request
	if err := req.ParseRequest(&request); err != nil {
		return contracts.Response{HttpStatus: http.StatusBadRequest, Err: err, Message: err.Error()}
	}

	resp, err := h.svc.Login(ctx, request)
	fmt.Println(resp, err)
	return contracts.Response{}
}
