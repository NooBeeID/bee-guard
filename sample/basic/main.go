package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/NooBeeID/bee-guard/beeguard"
	beeguardcache "github.com/NooBeeID/bee-guard/beeguard/cache"
)

func main() {
	router := http.NewServeMux()
	beeGuard := beeguard.New().
		SetRouter(router).
		SetDefaultModules().
		SetCache(beeguardcache.NewMemoryCache())

	if err := beeGuard.Run(); err != nil {
		log.Fatal(err)
	}

	slog.Info("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
