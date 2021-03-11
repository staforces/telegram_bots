package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	ExchangeID    string
	ClientID      string
	ClientSecret  string
	ServerHost    string
	ServerPort    string
	PgHost        string
	PgPort        string
	PgUser        string
	PgPass        string
	PgBase        string
	TelegramToken string
}

var cfg config

func init() {
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		log.Fatal(err)
	}

}
