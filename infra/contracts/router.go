package contracts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	body   []byte
	Query  map[string]any
	Params map[string]any
}

func (r *Request) SetBody(body []byte) *Request {
	r.body = body
	return r
}

func (r Request) ParseRequest(body any) error {
	err := json.Unmarshal(r.body, body)
	return err
}

type Response struct {
	HttpStatus     int    `json:"-"`
	Success        bool   `json:"success"`
	StatusCode     string `json:"status_code"`
	Message        string `json:"message"`
	Payload        any    `json:"payload"`
	Err            error  `json:"error"`
	AdditionalInfo any    `json:"additional_info"`
}

type opts func(r *Response) *Response

func WithError(err error) opts {
	return func(r *Response) *Response {
		r.Err = err
		r.AdditionalInfo = map[string]any{
			"error": err.Error(),
		}
		return r
	}
}

func WithPayload(payload interface{}) opts {
	return func(r *Response) *Response {
		r.Payload = payload
		return r
	}
}

func WithStatusCode(statusCode string) opts {
	return func(r *Response) *Response {
		r.StatusCode = fmt.Sprintf("%d%s", r.HttpStatus, statusCode)
		return r
	}
}

func WithAdditionalInfo(additionalInfo interface{}) opts {
	return func(r *Response) *Response {
		r.AdditionalInfo = additionalInfo
		return r
	}
}

func (r *Response) WithAdditionalInfo(additionalInfo interface{}) *Response {
	r.AdditionalInfo = additionalInfo
	return r
}

func (r *Response) WithAdditionalInfoError(err string) *Response {
	r.AdditionalInfo = map[string]string{
		"error": err,
	}
	return r
}

func generateResponse(success bool, message string, httpStatus int, opts ...opts) *Response {
	var Response = &Response{
		HttpStatus: httpStatus,
		Success:    success,
		Message:    message,
	}

	for _, opt := range opts {
		Response = opt(Response)
	}

	return Response
}

func NewSuccessOK(message string, opts ...opts) *Response {
	return generateResponse(true, message, http.StatusOK, opts...)
}

func NewSuccessCreated(message string, opts ...opts) *Response {
	return generateResponse(true, message, http.StatusCreated, opts...)
}

func NewErrBadRequest(err error, opts ...opts) *Response {
	allOpts := append(opts, WithError(err))
	return generateResponse(false, err.Error(), http.StatusBadRequest, allOpts...)
}
func NewErrUnauthorized(err error, opts ...opts) *Response {
	allOpts := append(opts, WithError(err))
	return generateResponse(false, err.Error(), http.StatusUnauthorized, allOpts...)
}

func NewErrNotFound(err error, opts ...opts) *Response {
	allOpts := append(opts, WithError(err))
	return generateResponse(false, err.Error(), http.StatusNotFound, allOpts...)
}

func NewErrConflict(err error, opts ...opts) *Response {
	allOpts := append(opts, WithError(err))
	return generateResponse(false, err.Error(), http.StatusConflict, allOpts...)
}

func NewErrUnprocessableEntity(err error, opts ...opts) *Response {
	allOpts := append(opts, WithError(err))
	return generateResponse(false, err.Error(), http.StatusUnprocessableEntity, allOpts...)
}

func NewErrGeneral(err error, opts ...opts) *Response {
	allOpts := append(opts, WithError(err))
	return generateResponse(false, err.Error(), http.StatusInternalServerError, allOpts...)
}

func (r Response) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return ""
}

func (r Response) IsError() bool {
	return r.Err != nil
}

type HandlerFunc func(ctx context.Context, req Request) Response

type Handler interface {
	Handle(ctx context.Context, req Request) Response
}