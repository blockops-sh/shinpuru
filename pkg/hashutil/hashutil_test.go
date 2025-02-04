package hashutil

import (
	"crypto"
	_ "crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	token = "my_secret_test_token"
)

func TestHashCompare(t *testing.T) {
	h := Hasher{
		HashFunc: crypto.SHA256,
		SaltSize: 128,
	}

	hash, err := h.Hash(token)
	assert.Nil(t, err)

	ok, err := Compare(token, hash)
	assert.Nil(t, err)
	assert.True(t, ok)

	// Lets flip some runes in the hash to make it invalid
	_hash := []byte(hash)
	if _hash[10] == 'a' {
		_hash[10] = 'b'
	} else {
		_hash[10] = 'a'
	}
	hash = string(_hash)

	ok, err = Compare(token, hash)
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestHashComparePepper(t *testing.T) {
	pepperGetter := func() ([]byte, error) {
		return []byte("my_pepper"), nil
	}

	h := Hasher{
		HashFunc:     crypto.SHA256,
		SaltSize:     128,
		PepperGetter: pepperGetter,
	}

	hash, err := h.Hash(token)
	assert.Nil(t, err)

	// First try without pepper, which must fail
	ok, err := Compare(token, hash)
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = Compare(token, hash, pepperGetter)
	assert.Nil(t, err)
	assert.True(t, ok)

	// Lets flip some runes in the hash to make it invalid
	_hash := []byte(hash)
	_hash[10] = 'a'
	_hash[11] = 'b'
	_hash[12] = 'c'
	hash = string(_hash)

	ok, err = Compare(token, hash, pepperGetter)
	assert.Nil(t, err)
	assert.False(t, ok)
}
