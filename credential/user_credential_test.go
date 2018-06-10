package credential_test

import (
	"reflect"
	"testing"

	"github.com/erichnascimento/nanosec/credential"
	"github.com/erichnascimento/nanosec/storage"
)

func TestCredentialWithCorrectPassword(t *testing.T) {
	t.Parallel()

	username, pwd := "erichnascimento", "pwd"
	uc, _ := createUserCredential(username)

	err := uc.SetPassword(pwd)
	if err != nil {
		t.Errorf(`Unexpected error when setting password: %v`, err)
	}

	match, err := uc.PasswordVerify(pwd)
	if err != nil {
		t.Errorf(`Unexpected error when verifying password: %v`, err)
	}

	if !match {
		t.Errorf(`Correct password does not matches on verification`)
	}

	newPwd := "newPwd"
	err = uc.ChangePassword(pwd, newPwd)
	if err != nil {
		t.Errorf(`Unexpected error when changing the password: %v`, err)
	}

	match, err = uc.PasswordVerify(newPwd)
	if err != nil {
		t.Errorf(`Unexpected error when verifying password: %v`, err)
	}
	if !match {
		t.Errorf(`Correct password does not matches on verification`)
	}
}

func TestCredentialWithIncorrectPassword(t *testing.T) {
	t.Parallel()

	username, pwd, otherPwd := "erichnascimento", "pwd", "otherPwd"
	uc, _ := createUserCredential(username)

	uc.SetPassword(pwd)
	match, _ := uc.PasswordVerify(otherPwd)

	if match {
		t.Errorf(`Incorrect password matches on verification`)
	}
}

func TestCredentialRoles(t *testing.T) {
	t.Parallel()

	username := "erichnascimento"
	uc, _ := createUserCredential(username)

	roles := []string{"admin", "buyer"}
	err := uc.AddRoles(roles...)
	if err != nil {
		t.Errorf(`Unexpected error when adding roles: %v`, err)
	}

	gaveRoles, err := uc.GetRoles()
	if err != nil {
		t.Errorf(`Unexpected error when getting roles: %v`, err)
	}

	if !reflect.DeepEqual(roles, gaveRoles) {
		t.Errorf(`Roles does not matches. expected = "%v", gave = "%v"`, roles, gaveRoles)
	}

	err = uc.RemoveRoles("admin", "buyer")
	if err != nil {
		t.Errorf(`Unexpected error when removing roles: %v`, err)
	}

	gaveRoles, err = uc.GetRoles()
	if err != nil {
		t.Errorf(`Unexpected error when getting roles after remove them: %v`, err)
	}

	if len(gaveRoles) != 0 {
		t.Errorf(`Roles was not removed. Expected a empty list, gave = %v`, gaveRoles)
	}
}

func createUserCredential(username string) (credential.UserCredential, error) {
	kvStorage, err := storage.NewMiniRedis()
	if err != nil {
		return nil, err
	}
	s := credential.NewKeyValueStorage(kvStorage)

	return credential.NewUserCredential(username, s), nil
}
