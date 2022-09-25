package rates

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

type mockCfg struct {
	ExchangeRate string `yaml:"exchange_rate"`
}

func NewMockCfg() (mockCfg, error) {
	cfg := mockCfg{}
	path, err := os.Getwd()
	if err != nil {
		return cfg, err
	}

	b, err := os.ReadFile(path + "/../../../cfg.yml")
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (c mockCfg) GetExchangeRateLink() string {
	return c.ExchangeRate
}

func TestUpdate(t *testing.T) {
	cfg, err := NewMockCfg()
	if err != nil {
		t.Error(err)
	}

	d := NewDailyRates(cfg)
	err = d.update()
	if err != nil {
		t.Error(err)
	}
	if len(d.Date) == 0 {
		t.Error("null map")
	}
}

func TestConvert(t *testing.T) {
	cfg, err := NewMockCfg()
	if err != nil {
		t.Error(err)
	}

	d := NewDailyRates(cfg)
	err = d.update()
	if err != nil {
		t.Error(err)
	}
	if len(d.Date) == 0 {
		t.Error("null map")
	}

	val, err := d.Exchange(65, "USD")
	if err != nil {
		t.Error(err)
	}
	if val < 0.8 || val > 1.2 {
		t.Error("Expected [0.8 usd < 65 rub < 1.2 usd]. Probably the exchange rates have changed")
	}
}

func TestConvertNegative(t *testing.T) {
	cfg, err := NewMockCfg()
	if err != nil {
		t.Error(err)
	}

	d := NewDailyRates(cfg)
	err = d.update()
	if err != nil {
		t.Error(err)
	}
	if len(d.Date) == 0 {
		t.Error("null map")
	}

	_, errRes := d.Exchange(65, "UStD")
	if errRes == nil {
		t.Errorf("Expected %s, but go ok status", errorWrongCurrency)
	}
}
