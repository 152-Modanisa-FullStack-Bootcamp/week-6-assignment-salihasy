package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

var c = &Config{}

func init() {
	file, err := os.Open(".config/local.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, c)
	if err != nil {
		panic(err)
	}
}
func Get() *Config {
	return c
}
