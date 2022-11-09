package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

// App это один большой класс(фасад), выполняющий всю работу
type App struct {
	history   History
	balances  Balances
	config    Config
	exchanger Exchanger

	mux *http.ServeMux
}

// NewApp это конструктор для класса App
func NewApp(history History, balances Balances, exchanger Exchanger, config Config) *App {
	a := &App{
		history:   history,
		balances:  balances,
		config:    config,
		exchanger: exchanger,
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

// ListenAndServe отвечает за запуск всех необходимых компонентов
func (a *App) ListenAndServe() error {
	log.Print("Starting interrupt listener")
	go a.interrupt()

	log.Print("Starting exchange updater")
	go a.exchanger.StartUpdater()

	log.Printf("Starting server at %s", a.config.GetHostPort())
	err := http.ListenAndServe(a.config.GetHostPort(), a.mux)
	return err
}

func (a *App) interrupt() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	for range sig {
		log.Println("Got interrupt signal")
	}
}
