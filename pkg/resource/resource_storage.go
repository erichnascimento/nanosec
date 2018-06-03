package resource

import (
	"fmt"

	"github.com/erichnascimento/nanosec/storage"
)

type Resource struct {
	Name string
}

type ResourceStorage interface {
	storage.Storage
	AddResource(string) (*Resource, error)
	RenameResource(string, string) (*Resource, error)
	GetResource(string) (*Resource, error)
	DeleteResource(string) error
}

var ResourceAlreadyExistsError = fmt.Errorf("Resource already exists")

type resourceStorage struct {
	storage.Storage
}

func (s *resourceStorage) AddResource(name string) (*Resource, error) {
	r := &Resource{name}
	if _, err := s.Get(name); err != storage.NotFoundError {
		return nil, ResourceAlreadyExistsError
	}
	s.Set(name, r)
	return r, nil
}

func (s *resourceStorage) RenameResource(oldName, newName string) (*Resource, error) {
	_, err := s.GetResource(oldName)
	if err != nil {
		return nil, err
	}

	s.DeleteResource(oldName)
	if err != nil {
		return nil, err
	}

	return s.AddResource(newName)
}

func (s *resourceStorage) GetResource(name string) (*Resource, error) {
	v, err := s.Get(name)
	if err != nil {
		return nil, err
	}

	r, ok := v.(*Resource)
	if !ok {
		return nil, fmt.Errorf(`Error retriving resource "%s". Expected Resource, given %v`, name, v)
	}

	return r, nil
}

func (s *resourceStorage) DeleteResource(name string) error {
	return s.Delete(name)
}

func NewResourceStorage(s storage.Storage) ResourceStorage {
	return &resourceStorage{s}
}
