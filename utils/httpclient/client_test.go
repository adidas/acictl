package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(123, 123, "they should be equals")
}
