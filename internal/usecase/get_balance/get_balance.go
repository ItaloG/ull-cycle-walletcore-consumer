package get_balance

import "github.com/ItaloG/full-cycle-walletcore-consumer/internal/gateway"

type GetBalanceInputDto struct {
	AccountID string `json:"account_id"`
}

type GetBalanceOutputDto struct {
	Balance float64
}

type GetBalanceUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewGetBalanceUseCase(accountGateway gateway.AccountGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		AccountGateway: accountGateway,
	}
}

func (uc *GetBalanceUseCase) Execute(input GetBalanceInputDto) (*GetBalanceOutputDto, error) {
	account, err := uc.AccountGateway.FindByID(input.AccountID)
	if err != nil {
		return nil, err
	}
	return &GetBalanceOutputDto{Balance: account.Balance}, nil
}
