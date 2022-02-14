package repository

import (
	"assignment/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetAll(t *testing.T) {

	repository := NewRepository()

	lengthOfWallet := len(repository.GetAll())
	assert.Equal(t, 2, lengthOfWallet)
}

func TestRepository_Create(t *testing.T) {
	wallet := model.Wallet{
		Username: "test",
		Balance:  11,
	}
	repository := NewRepository()

	initialLengthOfWallet := len(repository.GetAll())
	repository.Create(wallet)
	assert.Equal(t, initialLengthOfWallet+1, len(repository.GetAll()))
}

func TestRepository_GetByUsername(t *testing.T) {
	expectedWallet := model.Wallet{
		Username: "test",
		Balance:  11,
	}

	repository := NewRepository()
	repository.Create(expectedWallet)

	actual, err := repository.GetByUsername(expectedWallet.Username)
	assert.Equal(t, &expectedWallet, actual)
	assert.Nil(t, err)

}
func TestRepository_GetByUsername_Should_Return_UserNotFound_When_UserNotExist(t *testing.T) {
	repository := NewRepository()

	actual, err := repository.GetByUsername("notexistuser")
	assert.Equal(t, "User not found! ", err.Error())
	assert.Nil(t, actual)

}

func TestRepository_Update(t *testing.T) {
	wallet := model.Wallet{
		Username: "test",
		Balance:  11,
	}
	repository := NewRepository()
	repository.Create(wallet)
	willUpdateWallet := model.Wallet{
		Username: "test",
		Balance:  12,
	}
	repository.Update(willUpdateWallet)
	actualWallet, _ := repository.GetByUsername(willUpdateWallet.Username)
	assert.Equal(t, &willUpdateWallet, actualWallet)
}
