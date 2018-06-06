package credential

import (
	"errors"
	"fmt"
)

type UserCredential interface {
	SetPassword(password string) error
	ChangePassword(currentPassword, newPassword string) error
	AddRoles(roles ...string) error
	GetRoles() ([]string, error)
	RemoveRoles(roles ...string) error
	PasswordVerify(password string) (bool, error)
}

type userCredential struct {
	storage  Storage
	username string
}

func (uc *userCredential) SetPassword(password string) error {
	encPwd, err := passwordHash(password)
	if err != nil {
		return fmt.Errorf(`Error encoding password hash: %v`, err)
	}

	err = uc.storage.SetEncryptedPassword(uc.username, encPwd)
	if err != nil {
		return fmt.Errorf(`Error persisting encrypted password: %v`, err)
	}

	return nil
}

func (uc *userCredential) ChangePassword(currentPassword, newPassword string) error {
	match, err := uc.PasswordVerify(currentPassword)
	if err != nil {
		return fmt.Errorf(`Error verifying current password: %v`, err)
	}
	if !match {
		return ErrCurrentPasswordDoesNotMatch
	}

	newEncryptedPassword, err := passwordHash(newPassword)
	if err != nil {
		return err
	}

	return uc.SetPassword(newEncryptedPassword)
}

func (uc *userCredential) AddRoles(roles ...string) error {
	return uc.storage.AddRoles(uc.username, roles...)
}

func (uc *userCredential) GetRoles() ([]string, error) {
	return uc.storage.GetRoles(uc.username)
}

func (uc *userCredential) RemoveRoles(roles ...string) error {
	return uc.storage.RemoveRoles(uc.username, roles...)
}

func (uc *userCredential) PasswordVerify(password string) (bool, error) {
	encPwd, err := uc.storage.GetEncryptedPassword(uc.username)
	if err != nil {
		return false, fmt.Errorf(`Error getting encrypted password from storage: %v`, err)
	}

	match, err := passwordVerify(password, encPwd)
	if err != nil {
		return false, fmt.Errorf(`Error verifying password: %v`, err)
	}

	return match, nil
}

func NewUserCredential(username string, storage Storage) UserCredential {
	return &userCredential{
		username: username,
		storage:  storage,
	}
}

func passwordHash(password string) (string, error) {
	return password, nil
}

func passwordVerify(password, encryptedPassword string) (bool, error) {
	match := password == encryptedPassword

	return match, nil
}

var ErrCurrentPasswordDoesNotMatch = errors.New(`Current password informed does not matches`)
