package auth

import (
	"errors"
)

type Authenticator interface {
	Authenticate(credential string) error
}

func NewAuthenticator(storage Storage, crypter Crypter) Authenticator {
	return &authenticator{
		storage: storage,
		crypter: crypter,
	}
}

type authenticator struct {
	storage Storage
	crypter Crypter
}

func (a *authenticator) Authenticate(authorization string) error {
	credential, err := parseAuthorization(authorization, a.crypter)
	if err != nil {
		return ErrNotAuthorized
	}

	hasAuthorization, err := a.storage.HasAuthorization(credential, authorization)
	if err != nil {
		return err
	}

	if !hasAuthorization {
		return ErrNotAuthorized
	}

	return nil
}

var ErrNotAuthorized = errors.New(`Not authorized`)
