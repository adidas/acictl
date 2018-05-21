package httpclient

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert := assert.New(t)

	c := NewClient(
		"https://api.github.com",
		"GET",
		"application/json",
		"",
		"",
	)

	err := c.Run(
		bytes.Buffer{},
		func(res *http.Response, b string) {
			assert.Equal(200, res.StatusCode, "StatusCodes should be equals")
		},
		false,
		false,
	)

	assert.Nil(err, "Error should be nil")
}
