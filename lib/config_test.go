package lib

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPair(t *testing.T) {
	// Dummy seed will provide numbers 82 & 88
	rand.Seed(1)
	config := Config{}

	config.Coins = []Coin{
		{"ETH", 83, "USD", 43.1},
		{"BTC", 17, "USD", 123.21},
	}

	coin := config.GetCoin()
	assert.Equal(t, "ETH-USD", coin.GetPair(), "ETH first time")
	assert.Equal(t, 43.1, coin.Amount, "Correct amount")

	coin = config.GetCoin()
	assert.Equal(t, "BTC-USD", coin.GetPair(), "BTC second time")
	assert.Equal(t, 123.21, coin.Amount, "Correct amount")
}
