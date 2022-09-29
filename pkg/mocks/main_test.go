package mocks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test for a helper function used for testing
func TestRoute(t *testing.T) {
	assert.Equal(t, "http://127.0.0.1:8081/foo/bar", Route("foo", "bar"))
	assert.Equal(t, "http://127.0.0.1:8081/foo", Route("foo"))
	assert.Equal(t, "http://127.0.0.1:8081", Route())
}
