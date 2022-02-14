package service

import (
	"assignment/mock"
	"assignment/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	initialBalanceAmount = 0
	minimumBalanceAmount = -100
)

func TestService_GetAll(t *testing.T) {
	repositoryReturn := map[string]*model.Wallet{
		"pinpon": {
			Username: "pinpon",
			Balance:  100,
		},
		"saliha": {
			Username: "Saliha",
			Balance:  16,
		},
	}

	expectedWallets := make([]*model.Wallet, 0)
	expectedWallets = append(expectedWallets, &model.Wallet{
		Username: "pinpon",
		Balance:  100,
	}, &model.Wallet{
		Username: "Saliha",
		Balance:  16,
	})

	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().
		GetAll().
		Return(repositoryReturn).
		Times(1)

	service := NewService(repository, 0, -100)
	actualWallets := service.GetAll()

	assert.Equal(t, &expectedWallets, &actualWallets)
}
func TestService_Get(t *testing.T) {
	repositoryReturn := &model.Wallet{
		Username: "pinpon",
		Balance:  15,
	}

	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().
		GetByUsername("pinpon").
		Return(repositoryReturn, nil).
		Times(1)

	service := NewService(repository, initialBalanceAmount, minimumBalanceAmount)
	actualWallet, err := service.Get("pinpon")
	assert.Equal(t, repositoryReturn, actualWallet)
	assert.Nil(t, err)
}
func TestService_Get_Should_Return_Error_If_UserNotExist(t *testing.T) {
	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().
		GetByUsername("pinpon").
		Return(nil, errors.New("User not found! ")).
		Times(1)

	service := NewService(repository, initialBalanceAmount, minimumBalanceAmount)
	actualWallet, err := service.Get("pinpon")
	assert.Equal(t, "User not found! ", err.Error())
	assert.Nil(t, actualWallet)
}
func TestService_Create(t *testing.T) {
	expectedWallet := model.Wallet{
		Username: "pinpon",
		Balance:  initialBalanceAmount,
	}
	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().GetByUsername("pinpon").Return(nil, errors.New("test")).Times(1)
	repository.EXPECT().Create(expectedWallet).Return(&expectedWallet).Times(1)
	service := NewService(repository, initialBalanceAmount, minimumBalanceAmount)
	createdWallet := service.Create("pinpon")

	assert.Equal(t, &expectedWallet, createdWallet)
}
func TestService_Create_If_Exist_Return_Existing_One(t *testing.T) {
	expectedWallet := model.Wallet{
		Username: "pinpon",
		Balance:  100,
	}
	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().GetByUsername("pinpon").Return(&expectedWallet, nil).Times(1)
	service := NewService(repository, initialBalanceAmount, minimumBalanceAmount)
	createdWallet := service.Create("pinpon")

	assert.Equal(t, &expectedWallet, createdWallet)
}
func TestService_Update(t *testing.T) {
	repositoryWallet := model.Wallet{
		Username: "pinpon",
		Balance:  100,
	}
	expectedWallet := model.Wallet{
		Username: "pinpon",
		Balance:  80,
	}
	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().GetByUsername(repositoryWallet.Username).Return(&repositoryWallet, nil).Times(1)
	repository.EXPECT().Update(expectedWallet).Return(&expectedWallet).Times(1)
	service := NewService(repository, initialBalanceAmount, minimumBalanceAmount)

	actualWallet, err := service.Update(expectedWallet.Username, -20)
	assert.Equal(t, &expectedWallet, actualWallet)
	assert.Nil(t, err)
}
func TestService_Update_Should_Return_Error_If_Less_Than_MinimumBalanceAmount(t *testing.T) {
	repositoryWallet := model.Wallet{
		Username: "pinpon",
		Balance:  100,
	}
	expectedWallet := model.Wallet{
		Username: "pinpon",
		Balance:  80,
	}
	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().GetByUsername(repositoryWallet.Username).Return(&repositoryWallet, nil).Times(1)
	service := NewService(repository, initialBalanceAmount, minimumBalanceAmount)

	actualWallet, err := service.Update(expectedWallet.Username, -300)
	assert.NotNil(t, err)
	assert.Nil(t, actualWallet)
}
