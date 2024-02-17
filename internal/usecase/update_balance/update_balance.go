package update_balance

import (
	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/entity"
	"github.com/ItaloG/full-cycle-walletcore-consumer/internal/gateway"
)

type UpdateBalanceInputDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type UpdateBalanceUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewUpdateBalanceUseCase(accountGateway gateway.AccountGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		AccountGateway: accountGateway,
	}
}

func (uc *UpdateBalanceUseCase) Execute(input UpdateBalanceInputDTO) error {
	account, err := uc.AccountGateway.FindByID(input.AccountID)
	if err != nil {
		newAccount := entity.NewAccount(input.AccountID, input.Balance)
		return uc.AccountGateway.Save(newAccount)
	}

	return uc.AccountGateway.UpdateBalance(account)
}
