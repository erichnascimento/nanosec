package role_test

import (
	"fmt"
	"testing"

	"github.com/erichnascimento/nanosec/pkg/role"
	"github.com/erichnascimento/nanosec/storage"
)

func TestAddNewRole(t *testing.T) {
	t.Parallel()

	rs, s := newStorages()

	const name = "foo"
	r, err := rs.AddRole(name)
	if r == nil {
		t.Error(`Expected add method returning a Role instance. nil given`)
	}

	if err != nil {
		t.Errorf(`Unexpected error when adding a new role: "%v"`, err)
	}

	v, err := s.Get(name)
	if err != nil {
		t.Errorf(`Role not persisted: "%s"`, name)
	}

	if r := v.(*role.Role); r.Name != name {
		t.Errorf(`Role not persisted successfully. Expected r.name = "%s", given "%s"`, name, r.Name)
	}
}

func TestAddDuplicatedRole(t *testing.T) {
	t.Parallel()

	rs, _ := newStorages()

	const name = "foo"
	rs.AddRole(name)
	_, err := rs.AddRole(name)

	if err != storage.DocumentAlreadyExistsError {
		t.Errorf(`Adding duplicated role should return error`)
	}
}

func TestGetAddedRole(t *testing.T) {
	t.Parallel()

	rs, s := newStorages()

	const name = "fooRole"
	rs.AddRole(name)

	r, err := rs.GetRole(name)
	if err != nil {
		t.Errorf(`Unexpected error when getting role: %v`, err)
	}

	v, _ := s.Get(name)
	if v := v.(*role.Role); v.Name != r.Name {
		t.Errorf(`Name attr mismatch. Storage value: '%s', RoleStorage value = '%s'`, v.Name, r.Name)
	}
}

func TestGetNotAddedRole(t *testing.T) {
	t.Parallel()

	rs, _ := newStorages()

	const name = "fooRole"
	rs.AddRole(name)

	_, err := rs.GetRole("barRole")
	if err != storage.DocumentNotFoundError {
		t.Errorf(`NotFoundError expected when role was not added, given = "%v"`, err)
	}
}

func TestRenameExistentRole(t *testing.T) {
	t.Parallel()

	rs, s := newStorages()

	const name = "fooRole"
	rs.AddRole(name)

	const newName = "barRole"
	r, err := rs.RenameRole(name, newName)
	if err != nil {
		t.Errorf(`Unexpected error renaming role: %v`, err)
	}
	if r == nil {
		t.Error(`Returned value after renaming is nil. Expected a Role instance`)
	}

	if r.Name != newName {
		t.Errorf(`Returned value after renaming a role is wrong. Expected: "%s", given: "%s"`, newName, r.Name)
	}

	v, _ := s.Get(newName)
	r = v.(*role.Role)
	if r.Name != newName {
		t.Errorf(`Persisted value after renaming a role is wrong. Expected: "%s", given: "%s"`, newName, r.Name)
	}
}

func TestRenameNoExistentRole(t *testing.T) {
	t.Parallel()

	rs, _ := newStorages()

	const oldName = "fooRole"
	const newName = "barRole"
	_, err := rs.RenameRole(oldName, newName)
	if err != storage.DocumentNotFoundError {
		t.Errorf(`NotFoundError expected when trying to rename a no existent role, given = "%v"`, err)
	}
}

func ExampleRoleStorage() {
	rs, _ := newStorages()

	r1, _ := rs.AddRole("root")
	fmt.Println("Added role:", r1.Name)

	r2, _ := rs.RenameRole(r1.Name, "admin")
	fmt.Println("Renamed role from:", r1.Name, "to", r2.Name)

	r3, _ := rs.GetRole(r2.Name)
	fmt.Println("Getting role:", r3.Name)

	err := rs.DeleteRole(r3.Name)
	if err == nil {
		fmt.Println("Role:", r3.Name, "deleted sucessfully")
	}

	// Output:
	// Added role: root
	// Renamed role from: root to admin
	// Getting role: admin
	// Role: admin deleted sucessfully
}

func newStorages() (role.RoleStorage, storage.Storage) {
	s := storage.NewKeyValueStorage()
	rs := role.NewRoleStorage(s)

	return rs, s
}
