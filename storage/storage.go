package storage

import "fmt"

type Storage interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Delete(string) error
	// TODO Lock()
}

var DocumentNotFoundError = fmt.Errorf("Document not found")
var DocumentAlreadyExistsError = fmt.Errorf("Document already exists")
