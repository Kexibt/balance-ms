package app

type Database interface {
}

type History interface {
	GetByUID(userID string) []string
	Add(userID string, amount int)
}

type Balances interface {
	GetByUID(userID string) uint
	Change(userID string, amount int) (uint, error)
	Transfer(fromUserID, toUserID string, amount int) (uint, uint, error)
}

type Config interface {
	GetHostPort() string
}
