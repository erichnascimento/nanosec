package permission

import (
	"fmt"
)

type Granter interface {
	Grant([]string) error
	Revoke([]string) error
}

const grantErrorFmt = `Error granting access to "%s" for %v. Reason: %v`
const revokeErrorFmt = `Error revoking access to "%s" for %v. Reason: %v`

type granter struct {
	resource string
	storage  Storage
}

func (g *granter) Grant(roles []string) error {
	err := g.storage.AddRoles(roles)
	if err != nil {
		return fmt.Errorf(grantErrorFmt, g.resource, roles, err)
	}
	return nil
}

func (g *granter) Revoke(roles []string) error {
	err := g.storage.RemoveRoles(roles)
	if err != nil {
		return fmt.Errorf(revokeErrorFmt, g.resource, roles, err)
	}
	return nil
}

func NewGranter(resource string, s Storage) Granter {
	return &granter{
		resource: resource,
		storage:  s,
	}
}
