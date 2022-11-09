package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Kexibt/balance-microservice/internal/pkg/app"
	"github.com/Kexibt/balance-microservice/internal/pkg/balance"
	"github.com/Kexibt/balance-microservice/internal/pkg/config"
	"github.com/Kexibt/balance-microservice/internal/pkg/database"
	"github.com/Kexibt/balance-microservice/internal/pkg/rates"
)

func main() {
	cfg := config.GetConfig()

	log.Print("Opening connection to PostgreSQL")
	db := database.NewDatabase(cfg)

	app := app.NewApp(nil, balance.NewBalances(db), rates.NewDailyRates(cfg), cfg) // todo: rename exchanger....

	notif := make(chan os.Signal, 1)
	signal.Notify(notif, os.Interrupt)

	go func() {
		err := app.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-notif
	defer db.Close()
}
