package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/NooBeeID/bee-guard/beeguard"
)

func main() {
	router := http.NewServeMux()
	beeGuard := beeguard.New().SetRouter(router).SetDefaultModules()

	wrappedMux := applyMiddleware(router, func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("Middleware")
			h.ServeHTTP(w, r)
		})
	})

	if err := beeGuard.Run(); err != nil {
		log.Fatal(err)
	}

	slog.Info("Server running on port 8080")
	http.ListenAndServe(":8080", wrappedMux)
}

func applyMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}
