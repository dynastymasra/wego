package wego_test

import (
	"testing"
	"wego"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := wego.NewClient("token")

	assert.NotNil(t, client)
}

func TestNewClientWithBackOff(t *testing.T) {
	client := wego.NewClientWithBackOff("token", wego.NewBackOff(10, -1))

	assert.NotNil(t, client)
}
