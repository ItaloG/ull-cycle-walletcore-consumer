package update_balance

import (
	"errors"
	"testing"

	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/entity"
	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldUpdateAccountWhenAccountAlreadyExists(t *testing.T) {
	account := entity.NewAccount("id", 100)
	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("FindByID", account.ID).Return(account, nil)
	accountMock.On("UpdateBalance", account).Return(nil)

	uc := NewUpdateBalanceUseCase(accountMock)

	inputDto := UpdateBalanceInputDTO{
		AccountID: account.ID,
		Balance:   account.Balance,
	}
	err := uc.Execute(inputDto)
	assert.Nil(t, err)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 1)
	accountMock.AssertNumberOfCalls(t, "UpdateBalance", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 0)
	accountMock.AssertCalled(t, "UpdateBalance", account)
}

func TestShouldSaveAccountWhenFindByIdReturnAccountNotFoundError(t *testing.T) {
	account := entity.NewAccount("id", 100)
	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("FindByID", account.ID).Return(account, errors.New("account not found"))
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewUpdateBalanceUseCase(accountMock)

	inputDto := UpdateBalanceInputDTO{
		AccountID: account.ID,
		Balance:   account.Balance,
	}
	err := uc.Execute(inputDto)
	assert.Nil(t, err)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 1)
	accountMock.AssertNumberOfCalls(t, "UpdateBalance", 0)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
