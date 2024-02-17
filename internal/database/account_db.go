package database

import (
	"database/sql"
	"errors"

	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account

	stmt, err := a.DB.Prepare("Select id, balance, created_at FROM accounts WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.ID,
		&account.Balance,
		&account.CreatedAt,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("account not found")
	case err != nil:
		return nil, err
	default:
		return &account, nil
	}
}

func (a *AccountDB) Save(account *entity.Account) error {
	stmt, err := a.DB.Prepare("INSERT INTO accounts (id, balance, created_at) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.ID, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountDB) UpdateBalance(account *entity.Account) error {
	stmt, err := a.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return err
	}
	return nil
}
