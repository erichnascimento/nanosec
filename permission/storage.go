package permission

type Storage interface {
	AddRoles(resource string, roles ...string) error
	RemoveRoles(resource string, roles ...string) error
	HasAnyRole(resource string, roles ...string) (bool, error)
}
