package service

import (
	"assignment/model"
	"assignment/repository"
	"errors"
	"fmt"
)

type IWalletService interface {
	GetAll() []*model.Wallet
	Get(username string) (*model.Wallet, error)
	Create(username string) *model.Wallet
	Update(username string, balance int) (*model.Wallet, error)
}

type service struct {
	repo                 repository.IRepository
	InitialBalanceAmount int
	MinimumBalanceAmount int
}

func NewService(repo repository.IRepository,
	initialBalanceAmount int,
	minimumBalanceAmount int) IWalletService {
	return &service{
		repo:                 repo,
		InitialBalanceAmount: initialBalanceAmount,
		MinimumBalanceAmount: minimumBalanceAmount,
	}

}

func (s service) GetAll() []*model.Wallet {
	walletsMap := s.repo.GetAll()
	wallets := make([]*model.Wallet, 0)

	for _, wallet := range walletsMap {
		wallets = append(wallets, wallet)
	}
	return wallets
}

func (s service) Get(username string) (*model.Wallet, error) {
	wallet, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s service) Create(username string) *model.Wallet {
	existingWallet, err := s.repo.GetByUsername(username)
	if err == nil {
		return existingWallet
	}
	wallet := model.Wallet{
		Username: username,
		Balance:  s.InitialBalanceAmount,
	}

	return s.repo.Create(wallet)
}

func (s service) Update(username string, balance int) (*model.Wallet, error) {
	existingWallet, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	balanceResult := existingWallet.Balance + (balance)

	if balanceResult < s.MinimumBalanceAmount {
		return nil, errors.New(fmt.Sprintf("User's balance can not be less then %d", s.MinimumBalanceAmount))
	}
	wallet := model.Wallet{
		Username: username,
		Balance:  balanceResult,
	}
	updatedWallet := s.repo.Update(wallet)
	return updatedWallet, nil
}
