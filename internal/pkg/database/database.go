package database

import (
	"errors"
	"log"
)

type Database struct {
	dbCRUD *DatabaseCRUD
}

func NewDatabase(cfg Config) *Database {
	return &Database{
		dbCRUD: NewDatabaseCRUD(cfg),
	}
}

func (d *Database) Close() error {
	return d.dbCRUD.Close()
}

func (d *Database) Get(userid string) (float64, error) {
	log.Printf("geting balance {%s}\n", userid)
	return d.dbCRUD.GetBalance(userid)
}

func (d *Database) Change(userid string, new_balance float64) error {
	if new_balance < 0 {
		return errors.New("balance couldn't be less than 0")
	}
	log.Printf("changing balance {%s}: {%f}\n", userid, new_balance)
	return d.dbCRUD.ChangeBalance(userid, new_balance)
}

func (d *Database) Create(userid string, new_balance float64) error {
	if new_balance < 0 {
		return errors.New("balance couldn't be less than 0")
	}
	log.Printf("creating balance {%s}: {%f}\n", userid, new_balance)
	return d.dbCRUD.CreateBalance(userid, new_balance)
}
