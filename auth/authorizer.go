package auth

type Authorizer interface {
	Authorize(credential string) (string, error)
	Revoke(authorization string) error
}

type authorizer struct {
	storage Storage
	crypter Crypter
}

func (a *authorizer) Authorize(credential string) (string, error) {
	authorization, err := createAuthorization(credential, a.crypter)
	if err != nil {
		return "", err
	}

	err = a.storage.AddAuthorization(credential, authorization)
	if err != nil {
		return "", err
	}

	return authorization, nil
}

func (a *authorizer) Revoke(authorization string) error {
	credential, err := parseAuthorization(authorization, a.crypter)
	if err != nil {
		return err
	}

	return a.storage.DeleteAuthorization(credential, authorization)
}

func NewAuthorizer(storage Storage, crypter Crypter) Authorizer {
	return &authorizer{
		storage: storage,
		crypter: crypter,
	}
}

func createAuthorization(credential string, crypter Crypter) (string, error) {
	return crypter.Encrypt(credential)
}

func parseAuthorization(authorization string, crypter Crypter) (string, error) {
	return crypter.Decrypt(authorization)
}
