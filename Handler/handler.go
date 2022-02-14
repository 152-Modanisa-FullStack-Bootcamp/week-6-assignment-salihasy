package handler

import (
	"assignment/model"
	"assignment/service"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var (
	getAll                  = regexp.MustCompile(`^[\/]*$`)
	getWalletByUserName     = regexp.MustCompile(`^\/[a-z]*$`)
	createWalletforUserName = regexp.MustCompile(`^\/[a-z]*$`)
	updateWalletBalance     = regexp.MustCompile(`^\/[a-z]*$`)
)

type IHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IWalletService
}

func NewHandler(service service.IWalletService) IHandler {
	return &Handler{service: service}
}

type UpdateBalanceRequest struct {
	Balance int `json:"balance"`
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	wallets := h.service.GetAll()
	response := model.Response{
		Message: "",
		Data:    wallets,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	username := h.findStringSubmatch(r.URL.Path)
	wallet, err := h.service.Get(username)

	response := model.Response{
		Message: "",
		Data:    nil,
	}

	if err != nil {
		if err.Error() == "User not found! " {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		response.Message = err.Error()
	} else {
		response.Data = wallet
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	username := h.findStringSubmatch(r.URL.Path)
	wallet := h.service.Create(username)
	response := model.Response{
		Message: "",
		Data:    wallet,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	username := h.findStringSubmatch(r.URL.Path)
	balanceRequest := UpdateBalanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&balanceRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedWallet, err := h.service.Update(username, balanceRequest.Balance)

	response := model.Response{
		Message: "",
		Data:    nil,
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = err.Error()

	} else {
		response.Data = updatedWallet
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)
	return

}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && getAll.MatchString(r.URL.Path):
		h.GetAll(w, r)
		return
	case r.Method == http.MethodGet && getWalletByUserName.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPut && createWalletforUserName.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodPost && updateWalletBalance.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	default:
		fmt.Println("istek karşılanamadı")
		return
	}
}

func (h *Handler) findStringSubmatch(path string) string {
	return strings.TrimPrefix(path, "/")
}
