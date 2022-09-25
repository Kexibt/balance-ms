package database

import "time"

// Config интерфейс для конфига
type Config interface {
	GetConnectionTimeout() time.Duration
	GetConnectionString() string
}
