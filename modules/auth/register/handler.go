package register

import (
	"context"

	"github.com/NooBeeID/bee-guard/infra/contracts"
)

type ContractService interface {
	Register(ctx context.Context, req Request) (Response, error)
}
type handler struct {
	svc ContractService
}

func NewHandler(svc ContractService) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) Handle(ctx context.Context, req contracts.Request) any {

	var request Request
	if err := req.ParseRequest(&request); err != nil {
		return contracts.NewErrBadRequest(err, contracts.WithError(err))
	}

	resp, err := h.svc.Register(ctx, request)
	if err != nil {
		return err
	}
	return contracts.Response{
		Payload: resp,
	}
}
