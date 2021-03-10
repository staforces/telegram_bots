package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	ClientID     string
	ClientSecret string
	ServerHost   string
	ServerPort   string
	PgHost       string
	PgPort       string
	PgUser       string
	PgPass       string
	PgBase       string
}

var cfg config

func init() {
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		log.Fatal(err)
	}

}
