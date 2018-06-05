package permission

import (
	"github.com/erichnascimento/nanosec/storage"
)

type keyValueStorage struct {
	kvStorage storage.KeyValueStorage
}

func (s *keyValueStorage) AddRoles(resource string, roles ...string) error {
	_, err := s.kvStorage.SetAdd(resource, roles...)
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

func (s *keyValueStorage) HasAnyRole(resource string, roles ...string) (bool, error) {
	for _, role := range roles {
		isMember, err := s.kvStorage.IsMember(resource, role)
		if err != nil {
			return false, err
		}
		if isMember {
			return true, nil
		}
	}
	return false, nil
}

func NewKeyValueStorage(kvStorage storage.KeyValueStorage) (Storage, error) {
	storage := &keyValueStorage{
		kvStorage: kvStorage,
	}

	return storage, nil
}

const errorCreatingNewKeyValueStorageFmt = `Error when creating new KeyValueStorage. Reason: %v`
