package main

import (
	"assignment/config"
	"assignment/handler"
	"assignment/repository"
	"assignment/service"
	"log"
	"net/http"
)

func main() {
	config := config.Get()
	repository := repository.NewRepository()
	service := service.NewService(repository, config.InitialBalanceAmount, config.MinimumBalanceAmount)
	handler := handler.NewHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
