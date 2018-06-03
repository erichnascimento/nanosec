package storage

import "fmt"

type Storage interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Delete(string) error
	// TODO Lock()
}

var NotFoundError = fmt.Errorf("Document not found")
