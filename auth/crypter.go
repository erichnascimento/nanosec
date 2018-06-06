package auth

import "errors"

type Crypter interface {
	Encrypt(text string) (string, error)
	Decrypt(text string) (string, error)
}

type fakeCrypter struct {
	salt string
}

func (c *fakeCrypter) Encrypt(text string) (string, error) {
	return c.salt + text, nil
}

func (c *fakeCrypter) Decrypt(text string) (string, error) {
	if text[:len(c.salt)] != c.salt {
		return "", errors.New(`Invalid encrypted text`)
	}
	return text[len(c.salt):], nil
}

func NewFakeCrypter(salt string) Crypter {
	return &fakeCrypter{
		salt: salt,
	}
}
