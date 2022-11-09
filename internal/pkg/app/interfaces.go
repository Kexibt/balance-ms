package app

// History интерфейс отоброжающий нереализованный функционал класса History
type History interface {
	GetByUID(userID string) []string
	Add(userID string, oldBalance, newBalance, amount float64)
}

// Balances интерфейс отоброжающий функционал класса, отвечающего за баланс пользователей
type Balances interface {
	GetByUID(userID string) float64
	Change(userID string, amount float64) (float64, error)
	Transfer(fromUserID, toUserID string, amount float64) (float64, float64, error)
}

// Config интерфейс отоброжающий функционал класса, отвечающего за кофигурацию
type Config interface {
	GetHostPort() string
	GetExchangeRateLink() string
}

// Exchanger интерфейс отоброжающий функционал класса, отвечающего за курс валют
type Exchanger interface {
	Exchange(amount float64, currency string) (float64, error)
	StartUpdater()
}
