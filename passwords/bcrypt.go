package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

func NewBcryptPasswords(cost int) *BcryptPasswords {
	return &BcryptPasswords{
		cost: cost,
	}
}

type BcryptPasswords struct {
	cost int
}

func (p *BcryptPasswords) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (p *BcryptPasswords) Compare(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
