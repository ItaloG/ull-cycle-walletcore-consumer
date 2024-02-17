package database

import (
	"database/sql"
	"testing"

	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table accounts (id varchar(255), balance int, created_at date)")
	s.accountDB = NewAccountDB(db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount("id", 100)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	account := entity.NewAccount("id", 100)
	stmt, err := s.db.Prepare("INSERT INTO accounts (id, balance, created_at) VALUES (?, ?, ?)")
	s.Nil(err)
	defer stmt.Close()
	stmt.Exec(account.ID, account.Balance, account.CreatedAt)

	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Balance, accountDB.Balance)
}
