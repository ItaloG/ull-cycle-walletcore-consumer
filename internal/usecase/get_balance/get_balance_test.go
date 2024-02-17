package get_balance

import (
	"testing"

	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/entity"
	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnBalanceOnSuccess(t *testing.T) {
	account := entity.NewAccount("id", 100)
	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("FindByID", account.ID).Return(account, nil)
	uc := NewGetBalanceUseCase(accountMock)

	input := GetBalanceInputDto{
		AccountID: "id",
	}
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, output.Balance, account.Balance)
}
