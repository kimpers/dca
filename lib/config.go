package lib

import (
	"io/ioutil"
	"log"
	"math/rand"

	yaml "gopkg.in/yaml.v2"
)

type Coin struct {
	Ticker       string  `yaml:"ticker"`
	Percentage   float64 `yaml:"percentage"`
	BaseCurrency string  `yaml:"baseCurrency"`
	Amount       string  `yaml:"amount"`
}

func (coin Coin) GetPair() string {
	return coin.Ticker + "-" + coin.BaseCurrency
}

type Config struct {
	Coinbase struct {
		Secret     string `yaml:"secret"`
		Key        string `yaml:"key"`
		Passphrase string `yaml:"passphrase"`
	}
	Schedule string `yaml:"schedule"`
	Coins    []Coin
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

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("Failed to read yaml config file. Error: %v", err)
	}
}
