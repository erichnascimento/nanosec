package credential

import (
	"fmt"

	"github.com/erichnascimento/nanosec/storage"
)

type keyValueStorage struct {
	kvStorage storage.KeyValueStorage
}

func (s *keyValueStorage) GetEncryptedPassword(username string) (string, error) {
	pwd, err := s.kvStorage.Get(getEncryptedPasswordKey(username))
	if err != nil {
		return "", fmt.Errorf(`Was not possible get the encrypted password for username: "%s". Reason: %v`, username, err)
	}

	return pwd, nil
}

func (s *keyValueStorage) SetEncryptedPassword(username, encyptedPassword string) error {
	err := s.kvStorage.Set(getEncryptedPasswordKey(username), encyptedPassword)
	if err != nil {
		fmt.Errorf(`Error saving encrypted password: %v`, err)
	}

	return nil
}

func (s *keyValueStorage) AddRoles(username string, roles ...string) error {
	_, err := s.kvStorage.SetAdd(getRolesKey(username), roles...)
	if err != nil {
		return err
	}

	return nil
}

func (s *keyValueStorage) RemoveRoles(resource string, roles ...string) error {
	_, err := s.kvStorage.SRem(resource, roles...)
	if err != nil {
		return err
	}

	return nil
}

func (s *keyValueStorage) GetRoles(username string) ([]string, error) {
	return s.kvStorage.Members(getRolesKey(username))
}

func getEncryptedPasswordKey(username string) string {
	return fmt.Sprintf(`%s:encryptedPassword`, username)
}

func getRolesKey(username string) string {
	return fmt.Sprintf(`%s:roles`, username)
}

func NewKeyValueStorage(kvStorage storage.KeyValueStorage) Storage {
	return &keyValueStorage{
		kvStorage: kvStorage,
	}
}
