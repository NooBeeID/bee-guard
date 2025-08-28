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

func (r *Router) buildRouter(path, method string, handler contracts.HandlerFunc) {
	method = strings.ToUpper(method)
	switch r := r.router.(type) {
	case *http.ServeMux:
		r.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.ToUpper(r.Method) != method {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}
			ctx := r.Context()
			// request body
			req := contracts.Request{}
			var reqBody json.RawMessage
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			req.Body = reqBody
			fmt.Println(string(reqBody))

			resp := handler(ctx, req)
			json.NewEncoder(w).Encode(resp)

		}))
	}
}
