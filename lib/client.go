package lib

import (
  "sort"
  "fmt"
  "log"
	krakenApi "github.com/beldur/kraken-go-api-client"
)

type Client struct {
  kraken *krakenApi.KrakenApi
	config   *Config
}

func NewClient(config *Config) *Client {
  kraken := krakenApi.New(config.Kraken.Key, config.Kraken.Secret)
	client := Client{
	  kraken,
		config,
	}

	return &client
}

func (client *Client) Buy() (error) {
  coin := client.config.GetCoin()

  ticker := "X" + coin.Ticker + "Z" + coin.BaseCurrency
  trades, err := client.kraken.Trades(ticker, 0)

  if err != nil {
    log.Fatal(err)
  }


  sort.Slice(trades.Trades, func(i, j int) bool {
    return trades.Trades[i].Time > trades.Trades[j].Time
  })

  latestTrade := trades.Trades[0]

  fmt.Println(latestTrade)

  // Use latest trade to figure out how much we should by to get the equivalent of configured amount 
  volume := coin.Amount / latestTrade.PriceFloat

  fmt.Println(volume)

  result, err := client.kraken.AddOrder(ticker, "buy", "market", fmt.Sprintf("%f", volume), map[string]string {})

  if err != nil {
    log.Fatal(err)
  }

  log.Println(result.Description.Order)

  return nil
}
