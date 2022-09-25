package rates

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	errorWrongCurrency = errors.New("неверный формат валюты")
)

// DailyRates отвечает перевод денег по курсу
type DailyRates struct {
	Date  map[string]float64 `json:"rates"`
	cfg   Config
	mutex sync.RWMutex
}

// Config интерфейс для конфига
type Config interface {
	GetExchangeRateLink() string
}

// NewDailyRates конструктор для DailyRates
func NewDailyRates(cfg Config) *DailyRates {
	return &DailyRates{cfg: cfg}
}

// Exchange обменивает по актуальному курсу
func (d *DailyRates) Exchange(amount float64, currency string) (float64, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	val, exists := d.Date[strings.ToUpper(currency)]
	if !exists {
		return 0, errorWrongCurrency
	}

	return amount * val, nil
}

// StartUpdater запускает цикл обновлений курса валют
func (d *DailyRates) StartUpdater() {
	d.mutex.Lock()
	err := d.update()
	if err != nil {
		log.Println(err)
	}
	d.mutex.Unlock()

	ticker := time.NewTicker(time.Hour * 12) // можно было вынести в конфиг время обновления
	for range ticker.C {
		d.mutex.Lock()
		err := d.update()
		if err != nil {
			log.Println(err)
		}
		d.mutex.Unlock()
	}
}

func (d *DailyRates) update() error {
	resp, err := http.Get(d.cfg.GetExchangeRateLink())
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, d)
	if err != nil {
		return err
	}
	return nil
}
