package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencyCents(t *testing.T) {
	cents, err := Currency("1.23").Cents()
	if assert.NoError(t, err) {
		assert.Equal(t, cents, 23)
	}
}

func TestCurrencyDollars(t *testing.T) {
	dollars, err := Currency("1.23").Dollars()
	if assert.NoError(t, err) {
		assert.Equal(t, dollars, 1)
	}
}
