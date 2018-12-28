package lib

import (
	coinbase "github.com/preichenberger/go-gdax"
)

type Client struct {
	coinbase *coinbase.Client
	config   *Config
}

func NewClient(config *Config) *Client {
	coinbase := coinbase.NewClient(config.Coinbase.Secret, config.Coinbase.Key, config.Coinbase.Passphrase)
	client := Client{
		coinbase,
		config,
	}

	return &client
}

func (client *Client) Buy() (coinbase.Order, error) {
	coin := client.config.GetCoin()
	pair := coin.Ticker + "-" + coin.BaseCurrency
	order := coinbase.Order{
		Type:      "market",
		Side:      "buy",
		ProductId: pair,
		Funds:     coin.Amount,
	}

	return client.coinbase.CreateOrder(&order)
}
