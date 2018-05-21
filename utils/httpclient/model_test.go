package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		Client{
			Url:         "adidas.com",
			Method:      "GET",
			ContentType: "application/json",
			Password:    "",
			User:        "",
		},
		NewClient(
			"adidas.com",
			"GET",
			"application/json",
			"",
			"",
		),
		"Models should be equals",
	)
}
