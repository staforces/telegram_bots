package main

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

type TokenPrice struct {
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp"`
	Price                int   `json:"price"`
}

var tprices int

func wowapi() {
	TPrice := TokenPrice{}
	conf := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     "https://us.battle.net/oauth/token",
	}
	client := conf.Client(context.Background())
	url := "https://eu.api.blizzard.com/data/wow/token/index?namespace=dynamic-eu&locale=ru_RU"
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&TPrice)
	tprices = TPrice.Price / 10000
	err = insertdb(tprices)
	if err != nil {
		log.Fatal(err)
	}
}
