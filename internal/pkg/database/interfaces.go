package database

import "time"

type Config interface {
	GetConnectionTimeout() time.Duration
	GetConnectionString() string
}
