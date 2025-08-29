package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/NooBeeID/bee-guard/infra/contracts"
)

type Router struct {
	router any
}

func New(r any) *Router {
	return &Router{router: r}
}

func (r *Router) Post(path string, handler contracts.HandlerFunc) {
	r.buildRouter(path, http.MethodPost, handler)
}

func (router *Router) buildRouter(path, method string, handler contracts.HandlerFunc) {
	method = strings.ToUpper(method)
	switch r := router.router.(type) {
	case *http.ServeMux:
		r.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.ToUpper(r.Method) != method {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}
			ctx := r.Context()
			req := contracts.Request{}

			if router.isNeedRequestBody(method) {
				req.SetBody(r.Body)
			}

			resp := handler(ctx, req)
			json.NewEncoder(w).Encode(resp)
		}))
	default:
		panic(fmt.Sprintf("unknown router type %T", r))
	}
}

func (r *Router) isNeedRequestBody(method string) bool {
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		return true
	default:
		return false
	}
}
