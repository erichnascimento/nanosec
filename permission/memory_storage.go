package permission

import (
	"github.com/erichnascimento/nanosec/storage"
)

type memoryStorage struct {
	resource  string
	kvStorage storage.KeyValueStorage
}

func (s *memoryStorage) AddRoles(roles []string) error {
	_, err := s.kvStorage.SetAdd(s.resource, roles...)
	if err != nil {
		return err
	}

	return nil
}

func (s *memoryStorage) RemoveRoles(roles []string) error {
	_, err := s.kvStorage.SRem(s.resource, roles...)
	if err != nil {
		return err
	}

	return nil
}

func (s *memoryStorage) HasAnyRole(roles []string) (bool, error) {
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

func NewMemoryStorage(resource string, kvStorage storage.KeyValueStorage) (Storage, error) {
	storage := &memoryStorage{
		resource:  resource,
		kvStorage: kvStorage,
	}

	return storage, nil
}

const errorCreatingNewMemoryStorageFmt = `Error when creating new MemoryStorage. Reason: %v`
