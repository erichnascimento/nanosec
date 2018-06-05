package permission

type Storage interface {
	AddRoles([]string) error
	RemoveRoles([]string) error
	HasAnyRole([]string) (bool, error)
}
