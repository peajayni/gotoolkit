package passwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcryptPasswordsSmoke(t *testing.T) {
	p := NewBcryptPasswords(12)
	password := "password"

	hashed, err := p.Hash(password)
	assert.Nil(t, err)

	result := p.Compare(hashed, password)
	assert.True(t, result)

	result = p.Compare(hashed, "other")
	assert.False(t, result)
}
