package auth

import "github.com/erichnascimento/nanosec/storage"

type Storage interface {
	AddAuthorization(credential, authorization string) error
	DeleteAuthorization(credential, authorization string) error
	HasAuthorization(credential, authorization string) (bool, error)
}

func NewKeyValueStorage(kvStorage storage.KeyValueStorage) (Storage, error) {
	s := &keyValueStorage{
		kvStorage: kvStorage,
	}

	return s, nil
}

func (s *keyValueStorage) AddAuthorization(credential, authorization string) error {
	_, err := s.kvStorage.SetAdd(credential, authorization)
	return err
}

func (s *keyValueStorage) DeleteAuthorization(credential, authorization string) error {
	_, err := s.kvStorage.SRem(credential, authorization)
	return err
}

func (s *keyValueStorage) HasAuthorization(credential, authorization string) (bool, error) {
	hasAuthorization, err := s.kvStorage.IsMember(credential, authorization)
	if err == storage.ErrKeyNotFound {
		return false, nil
	}

	return hasAuthorization, err
}

type keyValueStorage struct {
	kvStorage storage.KeyValueStorage
}
