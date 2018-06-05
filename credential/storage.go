package credential

type Storage interface {
	GetEncryptedPassword(username string) (string, error)
	SetEncryptedPassword(username, encyptedPassword string) error
	AddRoles(username string, roles ...string) error
	RemoveRoles(username string, roles ...string) error
	GetRoles(username string) ([]string, error)
}
