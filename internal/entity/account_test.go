package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	account := NewAccount("id", 100)
	assert.NotNil(t, account)
}

func TestUpdateAccountBalance(t *testing.T) {
	account := NewAccount("id", 100)
	account.UpdateBalance(200)

	assert.Equal(t, account.Balance, float64(200))
}
