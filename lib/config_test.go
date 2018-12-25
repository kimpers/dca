package lib

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Coin struct {
	Ticker string
}

func TestGetCoin(t *testing.T) {
	// Dummy seed will provide numbers 82 & 88
	rand.Seed(1)
	config := Config{}

	config.Coins = []struct {
		Ticker     string  `yaml:"ticker"`
		Percentage float64 `yaml:"percentage"`
	}{
		{"ETH", 83},
		{"BTC", 17},
	}

	assert.Equal(t, "ETH", config.GetCoin(), "ETH first time")
	assert.Equal(t, "BTC", config.GetCoin(), "BTC second time")
}
