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

func (client *Client) Buy(coin string) (coinbase.Order, error) {
	pair := coin + "-" + client.config.Coinbase.BaseCurrency
	order := coinbase.Order{
		Type:      "market",
		Side:      "buy",
		ProductId: pair,
		Funds:     client.config.Amount,
	}

	return client.coinbase.CreateOrder(&order)
}
