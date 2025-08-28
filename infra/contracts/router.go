package contracts

import (
	"context"
	"encoding/json"
)

type Request struct {
	Body   json.RawMessage
	Query  map[string]any
	Params map[string]any
}

type Response struct {
	HttpStatus int
	Message    string
	Payload    any
	Err        error
}

func (r Response) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return r.Message
}

func (r Request) ParseRequest(body any) error {
	err := json.Unmarshal(r.Body, &body)
	return err
}

type HandlerFunc func(ctx context.Context, req Request) Response
