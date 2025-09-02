package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/NooBeeID/bee-guard/beeguard"
	"github.com/NooBeeID/bee-guard/infra/contracts"
	"github.com/NooBeeID/bee-guard/infra/modules"
	"github.com/NooBeeID/bee-guard/modules/auth/login"
)

func main() {
	router := http.NewServeMux()
	beeGuard := beeguard.New().SetRouter(router)
	beeGuard.SetCustomModules(login.New(&modules.ConfigService{
		Handler: CustomModuleLogin{},
	}))

	if err := beeGuard.Run(); err != nil {
		log.Fatal(err)
	}

	slog.Info("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

// sample custom modules
type CustomModuleLogin struct {
	svc login.ContractService
}

func (c CustomModuleLogin) Handle(ctx context.Context, req contracts.Request) any {
	fmt.Println("Custom Handle")
	return contracts.Response{}
}
