package main

import (
	"log"

	"github.com/Kexibt/balance-microservice/internal/pkg/app"
	"github.com/Kexibt/balance-microservice/internal/pkg/balance"
	"github.com/Kexibt/balance-microservice/internal/pkg/config"
)

func main() {
	cfg := config.GetConfig()
	app := app.NewApp(nil, nil, balance.NewBalances(), cfg)
	err := app.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
