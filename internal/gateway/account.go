package gateway

import "github.com/ItaloG/full-cycle-walletcore-consumer/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
