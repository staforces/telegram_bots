package main

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp"`
	Price                int    `json:"price"`
	ClientID             string `json:"ClientID"`
	ClientSecret         string `json:"ClientSecret"`
	ServerHost           string `json:"ServerHost"`
	ServerPort           string `json:"ServerPort"`
	PgHost               string `json:"PgHost"`
	PgPort               string `json:"PgPort"`
	PgUser               string `json:"PgUser"`
	PgPass               string `json:"PgPass"`
	PgBase               string `json:"PgBase"`
}

var cfg config

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
		panic("Не удалось открыть файл")
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&cfg)
}
