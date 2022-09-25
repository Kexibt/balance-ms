package main

import (
	"log"

	"github.com/Kexibt/balance-microservice/internal/pkg/app"
	"github.com/Kexibt/balance-microservice/internal/pkg/balance"
	"github.com/Kexibt/balance-microservice/internal/pkg/config"
	"github.com/Kexibt/balance-microservice/internal/pkg/database"
	"github.com/Kexibt/balance-microservice/internal/pkg/rates"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewDatabase(cfg)

	app := app.NewApp(db, nil, balance.NewBalances(), rates.NewDailyRates(cfg), cfg) // todo: rename exchanger....
	err := app.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
