package login

import (
	"context"
	"fmt"
)

type contractRepository interface {
}
type service struct {
	repo contractRepository
}

// Login implements contractService.
func (s service) Login(ctx context.Context, req Request) (Response, error) {
	fmt.Println("Login Service")
	return Response{}, nil
}

func NewService(repo contractRepository) service {
	return service{repo: repo}
}
