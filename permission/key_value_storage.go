package permission

import (
	"github.com/erichnascimento/nanosec/storage"
)

type keyValueStorage struct {
	resource  string
	kvStorage storage.KeyValueStorage
}

func (s *keyValueStorage) AddRoles(roles []string) error {
	_, err := s.kvStorage.SetAdd(s.resource, roles...)
	if err != nil {
		return err
	}

	return nil
}

func (s *keyValueStorage) RemoveRoles(roles []string) error {
	_, err := s.kvStorage.SRem(s.resource, roles...)
	if err != nil {
		return err
	}

	return nil
}

func (s *keyValueStorage) HasAnyRole(roles []string) (bool, error) {
	for _, role := range roles {
		isMember, err := s.kvStorage.IsMember(s.resource, role)
		if err != nil {
			return false, err
		}
		if isMember {
			return true, nil
		}
	}
	return false, nil
}

func NewKeyValueStorage(resource string, kvStorage storage.KeyValueStorage) (Storage, error) {
	storage := &keyValueStorage{
		resource:  resource,
		kvStorage: kvStorage,
	}

	return storage, nil
}

const errorCreatingNewKeyValueStorageFmt = `Error when creating new KeyValueStorage. Reason: %v`
