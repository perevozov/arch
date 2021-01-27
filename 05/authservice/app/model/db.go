package model

import "github.com/jmoiron/sqlx"

type transactionedFunc func() error

type Datastore interface {
	AddUser(user *User) (int64, error)
	UpdateUser(user *User) error
	SetUserPassword(user *User, password string) error
	DeleteUser(userID int64) error
	LoadUserWithId(userId int64) (*User, error)
	CheckUserPassword(username string, password string) (*User, error)
	WithTransaction(transactionedFunc) error
}

type DB struct {
	*sqlx.DB
}

func (db *DB) WithTransaction(cb transactionedFunc) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	callbackErr := cb()

	if callbackErr != nil {
		tx.Rollback()
		return callbackErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return commitErr
	}
	return nil
}
