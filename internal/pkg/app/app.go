package app

import (
	"log"
	"net/http"
)

type App struct {
	db       Database
	history  History
	balances Balances
	config   Config

	mux *http.ServeMux
}

func NewApp(db Database, history History, balances Balances, config Config) *App {
	a := &App{
		db:       db,
		history:  history,
		balances: balances,
		config:   config,
	}
	a.newServeMux()
	return a
}

func (a *App) newServeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.mainHandler)
	mux.HandleFunc("/get_balance", a.getBalanceHandler)
	mux.HandleFunc("/add_to_balance", a.addBalanceHandler)
	mux.HandleFunc("/transfer_balance", a.transferBalanceHandler)
	a.mux = mux
}

func (a *App) ListenAndServe() error {
	log.Printf("Starting server at %s", a.config.GetHostPort())
	err := http.ListenAndServe(a.config.GetHostPort(), a.mux)
	return err
}
