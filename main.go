package main

import (
	"log"

	"github.com/Kexibt/balance-microservice/internal/pkg/app"
	"github.com/Kexibt/balance-microservice/internal/pkg/balance"
	"github.com/Kexibt/balance-microservice/internal/pkg/config"
	"github.com/Kexibt/balance-microservice/internal/pkg/database"
)

func main() {
	cfg := config.GetConfig()
	app := app.NewApp(database.NewDatabase(cfg), nil, balance.NewBalances(), cfg)
	err := app.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
