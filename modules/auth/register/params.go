package register

import (
	"github.com/NooBeeID/bee-guard/entity"
	"github.com/google/uuid"
)

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r Request) toAuth() entity.Auth {
	return entity.Auth{
		ID:       uuid.NewString(),
		Email:    r.Email,
		Password: r.Password,
	}
}

type Response struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}
