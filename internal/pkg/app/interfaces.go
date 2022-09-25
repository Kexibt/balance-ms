package app

type Database interface {
	Close() error
	Add(map[string]interface{})
	ListenAndServe()
}

type History interface {
	GetByUID(userID string) []string
	Add(userID string, oldBalance, newBalance, amount float64)
}

type Balances interface {
	GetByUID(userID string) float64
	Change(userID string, amount float64) (float64, error)
	Transfer(fromUserID, toUserID string, amount float64) (float64, float64, error)
}

type Config interface {
	GetHostPort() string
	GetExchangeRateLink() string
}

type Exchanger interface {
	Exchange(amount float64, currency string) (float64, error)
	StartUpdater()
}
