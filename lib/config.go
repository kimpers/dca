package lib

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

type Coin struct {
	Ticker               string  `json:"ticker"`
	Percentage           float64 `json:"percentage"`
	BaseCurrency         string  `json:"baseCurrency"`
	AmountInBaseCurrency float64 `json:"amountInBaseCurrency"`
}

func (coin Coin) GetPair() string {
	return coin.Ticker + "-" + coin.BaseCurrency
}

type Config struct {
	Token  string `json:token`
	Kraken struct {
		Secret string `json:"secret"`
		Key    string `json:"key"`
	}
	Coins []Coin
}

func (config Config) GetCoin() Coin {
	r := float64(rand.Intn(100) + 1)

	var max float64
	var current Coin
	for _, coin := range config.Coins {
		if max > r {
			break
		}

		max += coin.Percentage
		current = coin
	}

	return current
}

func (config *Config) ReadFile(path string) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Could not read config file. Error: %v", err)
	}

	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("Failed to read yaml config file. Error: %v", err)
	}
}

func (config *Config) DecodeBase64Config(encoded string) {
	decoded, err := base64.URLEncoding.DecodeString(encoded)

	if err != nil {
		log.Fatalf("Failed to decode encoded config. Error: %v", err)
	}

	err = json.Unmarshal(decoded, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal decoded config. Error: %v", err)
	}
}

func (config *Config) EncodeBase64ConfigFile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Could not read config file. Error: %v", err)
	}

	encoded := base64.URLEncoding.EncodeToString(data)

	return encoded
}
