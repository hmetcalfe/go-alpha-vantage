package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeAPIKey(t *testing.T) {
	key := "Zm9vYmFy"
	expected := "foobar"

	result, err := DecodeAPIKey(key)
	assert.NoError(t, err)

	assert.Equal(t, expected, result)

	_, err = DecodeAPIKey("")
	assert.Error(t, err)

	_, err = DecodeAPIKey("%%%")
	assert.Error(t, err)
}

func TestEncodeAPIKey(t *testing.T) {
	expected := "Zm9vYmFy"
	key := "foobar"

	result := EncodeAPIKey(key)

	assert.Equal(t, expected, result)
}
