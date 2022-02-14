package repository

import (
	"assignment/model"
	"errors"
)

type IRepository interface {
	GetAll() map[string]*model.Wallet
	GetByUsername(username string) (*model.Wallet, error)
	Create(wallet model.Wallet) *model.Wallet
	Update(wallet model.Wallet) *model.Wallet
}

type Repository struct {
	Wallet map[string]*model.Wallet
}

func NewRepository() IRepository {
	return &Repository{
		Wallet: map[string]*model.Wallet{
			"pinpon": {
				Username: "pinpon",
				Balance:  100,
			},
			"saliha": {
				Username: "Saliha",
				Balance:  16,
			},
		},
	}
}

func (r *Repository) Update(wallet model.Wallet) *model.Wallet {
	r.Wallet[wallet.Username] = &model.Wallet{
		Username: wallet.Username,
		Balance:  wallet.Balance,
	}
	return r.Wallet[wallet.Username]
}

func (r *Repository) GetAll() map[string]*model.Wallet {
	a := r.Wallet
	return a
}
func (r *Repository) GetByUsername(username string) (*model.Wallet, error) {
	value, exist := r.Wallet[username]
	if !exist {
		return nil, errors.New("User not found! ")
	}
	return value, nil
}

func (r *Repository) Create(wallet model.Wallet) *model.Wallet {
	r.Wallet[wallet.Username] = &model.Wallet{
		Username: wallet.Username,
		Balance:  wallet.Balance,
	}
	return r.Wallet[wallet.Username]
}
