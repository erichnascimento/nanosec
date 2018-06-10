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

	return pwd.(string), nil
}

func (s *keyValueStorage) SetEncryptedPassword(username, encyptedPassword string) error {
	err := s.kvStorage.Set(getEncryptedPasswordKey(username), encyptedPassword)
	if err != nil {
		return fmt.Errorf(`Error saving encrypted password: %v`, err)
	}

	return nil
}

func (s *keyValueStorage) AddRoles(username string, roles ...string) error {
	_, err := s.kvStorage.SetAdd(getRolesKey(username), storage.StrListToInterfaceList(roles)...)
	if err != nil {
		return err
	}

	return nil
}

func (s *keyValueStorage) RemoveRoles(username string, roles ...string) error {
	_, err := s.kvStorage.SRem(getRolesKey(username), storage.StrListToInterfaceList(roles)...)
	if err == nil || err == storage.ErrKeyNotFound {
		return nil
	}

	return err
}

func (s *keyValueStorage) GetRoles(username string) ([]string, error) {
	roles, err := s.kvStorage.Members(getRolesKey(username))
	if err == nil || err == storage.ErrKeyNotFound {
		return storage.InterfaceListToStrList(roles), nil
	}

	return nil, err
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
