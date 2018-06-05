package permission

import (
	"fmt"
)

const hasAccessErrorFmt = `Error when checking if "%v" has access to "%s". Reason: %v`

type Checker interface {
	HasAccess([]string) (bool, error)
}

type checker struct {
	resource string
	s        Storage
}

func (c *checker) HasAccess(roles []string) (bool, error) {
	hasAccess, err := c.s.HasAnyRole(c.resource, roles...)
	if err != nil {
		err = fmt.Errorf(hasAccessErrorFmt, roles, c.resource, err)
	}

	return hasAccess, err
}

func NewChecker(resource string, s Storage) Checker {
	return &checker{
		resource: resource,
		s:        s,
	}
}
