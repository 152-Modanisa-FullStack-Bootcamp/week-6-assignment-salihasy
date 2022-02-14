package handler

import (
	"assignment/mock"
	"assignment/model"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestResponseForList struct {
	Message string          `json:"message"`
	Data    []*model.Wallet `json:"data"`
}

type TestResponse struct {
	Message string        `json:"message"`
	Data    *model.Wallet `json:"data"`
}

func TestHandler_GetAll(t *testing.T) {
	expectedWallets := make([]*model.Wallet, 0)
	expectedWallets = append(expectedWallets, &model.Wallet{
		Username: "pinpon",
		Balance:  100,
	}, &model.Wallet{
		Username: "Saliha",
		Balance:  16,
	})
	expectedResp := TestResponseForList{
		Message: "",
		Data:    expectedWallets,
	}
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().
		GetAll().
		Return(expectedWallets).
		Times(1)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	actual := TestResponseForList{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, &expectedResp, &actual)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestHandler_Get(t *testing.T) {
	wallet := model.Wallet{
		Username: "pinpon",
		Balance:  16,
	}
	expectedResp := TestResponse{
		Message: "",
		Data:    &wallet,
	}
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().
		Get("pinpon").
		Return(&wallet, nil).
		Times(1)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/pinpon", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	actual := TestResponse{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, &expectedResp, &actual)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestHandler_Create(t *testing.T) {
	wallet := model.Wallet{
		Username: "test",
		Balance:  0,
	}
	expectedResp := TestResponse{
		Message: "",
		Data:    &wallet,
	}
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().
		Create("test").
		Return(&wallet).
		Times(1)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodPut, "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	actual := TestResponse{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, &expectedResp, &actual)
	assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestHandler_Update(t *testing.T) {
	wallet := model.Wallet{
		Username: "test",
		Balance:  100,
	}
	expectedResp := TestResponse{
		Message: "",
		Data:    &wallet,
	}
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().
		Update("test", 100).
		Return(&wallet, nil).
		Times(1)

	handler := NewHandler(service)
	var jsonStr = []byte(`{"balance":100}`)
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	actual := TestResponse{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, &expectedResp, &actual)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}
