package role

import (
	"fmt"

	"github.com/erichnascimento/nanosec/storage"
)

type Role struct {
	Name string
}

type RoleStorage interface {
	AddRole(string) (*Role, error)
	RenameRole(string, string) (*Role, error)
	GetRole(string) (*Role, error)
	DeleteRole(string) error
}

type roleStorage struct {
	storage.Storage
}

func (s *roleStorage) AddRole(name string) (*Role, error) {
	r := &Role{name}
	if _, err := s.Get(name); err != storage.DocumentNotFoundError {
		return nil, storage.DocumentAlreadyExistsError
	}
	s.Set(name, r)
	return r, nil
}

func (s *roleStorage) RenameRole(oldName, newName string) (*Role, error) {
	_, err := s.GetRole(oldName)
	if err != nil {
		return nil, err
	}

	s.DeleteRole(oldName)
	if err != nil {
		return nil, err
	}

	return s.AddRole(newName)
}

func (s *roleStorage) GetRole(name string) (*Role, error) {
	v, err := s.Get(name)
	if err != nil {
		return nil, err
	}

	r, ok := v.(*Role)
	if !ok {
		return nil, fmt.Errorf(`Error retriving role "%s". Expected Role, given %v`, name, v)
	}

	return r, nil
}

func (s *roleStorage) DeleteRole(name string) error {
	return s.Delete(name)
}

func NewRoleStorage(s storage.Storage) RoleStorage {
	return &roleStorage{s}
}
