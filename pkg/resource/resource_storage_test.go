package resource_test

import (
	"fmt"
	"testing"

	"github.com/erichnascimento/nanosec/pkg/resource"
	"github.com/erichnascimento/nanosec/storage"
)

func TestAddNewResource(t *testing.T) {
	t.Parallel()

	rs, s := newStorages()

	const name = "foo"
	r, err := rs.AddResource(name)
	if r == nil {
		t.Error(`Expected add method returning a Resource instance. nil given`)
	}

	if err != nil {
		t.Errorf(`Unexpected error when adding a new resource: "%v"`, err)
	}

	v, err := s.Get(name)
	if err != nil {
		t.Errorf(`Resource not persisted: "%s"`, name)
	}

	if r := v.(*resource.Resource); r.Name != name {
		t.Errorf(`Resource not persisted successfully. Expected r.name = "%s", given "%s"`, name, r.Name)
	}
}

func TestAddDuplicatedResource(t *testing.T) {
	t.Parallel()

	rs, _ := newStorages()

	const name = "foo"
	rs.AddResource(name)
	_, err := rs.AddResource(name)

	if err != resource.ResourceAlreadyExistsError {
		t.Errorf(`Adding duplicated resource should return error`)
	}
}

func TestGetAddedResource(t *testing.T) {
	t.Parallel()

	rs, s := newStorages()

	const name = "fooResource"
	rs.AddResource(name)

	r, err := rs.GetResource(name)
	if err != nil {
		t.Errorf(`Unexpected error when getting resource: %v`, err)
	}

	v, _ := s.Get(name)
	if v := v.(*resource.Resource); v.Name != r.Name {
		t.Errorf(`Name attr mismatch. Storage value: '%s', ResourceStorage value = '%s'`, v.Name, r.Name)
	}
}

func TestGetNotAddedResource(t *testing.T) {
	t.Parallel()

	rs, _ := newStorages()

	const name = "fooResource"
	rs.AddResource(name)

	_, err := rs.GetResource("barResource")
	if err != storage.NotFoundError {
		t.Errorf(`NotFoundError expected when resource was not added, given = "%v"`, err)
	}
}

func TestRenameExistentResource(t *testing.T) {
	t.Parallel()

	rs, s := newStorages()

	const name = "fooResource"
	rs.AddResource(name)

	const newName = "barResource"
	r, err := rs.RenameResource(name, newName)
	if err != nil {
		t.Errorf(`Unexpected error renaming resource: %v`, err)
	}
	if r == nil {
		t.Error(`Returned value after renaming is nil. Expected a Resource instance`)
	}

	if r.Name != newName {
		t.Errorf(`Returned value after renaming a resource is wrong. Expected: "%s", given: "%s"`, newName, r.Name)
	}

	v, _ := s.Get(newName)
	r = v.(*resource.Resource)
	if r.Name != newName {
		t.Errorf(`Persisted value after renaming a resource is wrong. Expected: "%s", given: "%s"`, newName, r.Name)
	}
}

func TestRenameNoExistentResource(t *testing.T) {
	t.Parallel()

	rs, _ := newStorages()

	const oldName = "fooResource"
	const newName = "barResource"
	_, err := rs.RenameResource(oldName, newName)
	if err != storage.NotFoundError {
		t.Errorf(`NotFoundError expected when trying to rename a no existent resource, given = "%v"`, err)
	}
}

func ExampleResourceStorage() {
	rs, _ := newStorages()

	r1, _ := rs.AddResource("org.corporation.add")
	fmt.Println("Added resource:", r1.Name)

	r2, _ := rs.RenameResource(r1.Name, "org.corporation.create")
	fmt.Println("Renamed resource from:", r1.Name, "to", r2.Name)

	r3, _ := rs.GetResource(r2.Name)
	fmt.Println("Getting resource:", r3.Name)

	err := rs.DeleteResource(r3.Name)
	if err == nil {
		fmt.Println("Resource:", r3.Name, "deleted sucessfully")
	}

	// Output:
	// Added resource: org.corporation.add
	// Renamed resource from: org.corporation.add to org.corporation.create
	// Getting resource: org.corporation.create
	// Resource: org.corporation.create deleted sucessfully
}

func newStorages() (resource.ResourceStorage, storage.Storage) {
	s := storage.NewKeyValueStorage()
	rs := resource.NewResourceStorage(s)

	return rs, s
}
