package auth_test

import (
	"testing"

	"github.com/erichnascimento/nanosec/auth"
	"github.com/erichnascimento/nanosec/storage"
)

func TestAuthenticationForAuthorizedAccess(t *testing.T) {
	t.Parallel()

	redis, _ := storage.NewMiniRedis()
	s, _ := auth.NewKeyValueStorage(redis)
	crypter := auth.NewFakeCrypter("salt")
	authorizer := auth.NewAuthorizer(s, crypter)
	authorization, _ := authorizer.Authorize("erichnascimento")

	authenticator := auth.NewAuthenticator(s, crypter)
	err := authenticator.Authenticate(authorization)
	if err != nil {
		t.Errorf(`Unexpected error for authorized access: %v`, err)
	}
}

func TestAuthenticationForUnathorizedAccess(t *testing.T) {
	t.Parallel()

	redis, _ := storage.NewMiniRedis()
	s, _ := auth.NewKeyValueStorage(redis)
	crypter := auth.NewFakeCrypter("salt")

	authenticator := auth.NewAuthenticator(s, crypter)
	err := authenticator.Authenticate("erichnascimento")
	if err != auth.ErrNotAuthorized {
		t.Errorf(`Unauthorized access should return NotAuthorized error, given: %v`, err)
	}
}

func TestAuthenticationForRevokedAccess(t *testing.T) {
	t.Parallel()

	redis, _ := storage.NewMiniRedis()
	s, _ := auth.NewKeyValueStorage(redis)
	crypter := auth.NewFakeCrypter("salt")
	authorizer := auth.NewAuthorizer(s, crypter)
	authorization, _ := authorizer.Authorize("erichnascimento")
	authorizer.Revoke(authorization)

	authenticator := auth.NewAuthenticator(s, crypter)
	err := authenticator.Authenticate(authorization)
	if err != auth.ErrNotAuthorized {
		t.Errorf(`Revoked access should return NotAuthorized error, given: %v`, err)
	}
}
