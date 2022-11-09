package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockConfig struct{}

func (m mockConfig) GetConnectionString() string {
	return "postgres://postgres:postgrespw@localhost:55000/micro_balance"
}

func (m mockConfig) GetConnectionTimeout() time.Duration {
	return time.Second
}

func TestChange(t *testing.T) {
	m := mockConfig{}
	db := NewDatabaseCRUD(m)

	err := db.ChangeBalance("user_test10", 0)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	m := mockConfig{}
	db := NewDatabaseCRUD(m)

	amount, err := db.GetBalance("user_test")
	assert.NoError(t, err)
	assert.Equal(t, float64(0), amount)
}

func TestCreate(t *testing.T) {
	m := mockConfig{}
	db := NewDatabaseCRUD(m)

	err := db.CreateBalance("user_test10", 1250)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	m := mockConfig{}
	db := NewDatabaseCRUD(m)

	err := db.DeleteBalance("user_test")
	assert.NoError(t, err)
}
