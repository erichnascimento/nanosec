package permission_test

import (
	"testing"

	"github.com/erichnascimento/nanosec/permission"
	"github.com/erichnascimento/nanosec/storage"
)

func TestNotGrantedAccess(t *testing.T) {
	t.Parallel()

	redis, _ := storage.NewMiniRedis()
	defer redis.Close()

	resource := "my.printer"
	storage, _ := permission.NewKeyValueStorage(redis)

	granter := permission.NewGranter(resource, storage)
	granter.Grant("root")

	checker := permission.NewChecker(resource, storage)
	hasAccess, _ := checker.HasAccess("admin")
	if hasAccess {
		t.Error(`Access allowed for not granted role`)
	}
}

func TestGrantedAccess(t *testing.T) {
	t.Parallel()

	redis, _ := storage.NewMiniRedis()
	defer redis.Close()

	resource := "my.printer"
	storage, _ := permission.NewKeyValueStorage(redis)

	granter := permission.NewGranter(resource, storage)
	granter.Grant("admin", "attendent")

	checker := permission.NewChecker(resource, storage)
	hasAccess, _ := checker.HasAccess("admin")
	if !hasAccess {
		t.Error(`Access denied for granted role`)
	}

	hasAccess, _ = checker.HasAccess("attendent")
	if !hasAccess {
		t.Error(`Access denied for granted role`)
	}
}

func TestRevokedAccess(t *testing.T) {
	t.Parallel()

	redis, _ := storage.NewMiniRedis()
	defer redis.Close()

	r := "my.printer"
	s, _ := permission.NewKeyValueStorage(redis)

	granter := permission.NewGranter(r, s)
	granter.Grant("admin")
	granter.Revoke("admin")

	checker := permission.NewChecker(r, s)
	hasAccess, _ := checker.HasAccess("admin")
	if hasAccess {
		t.Error(`Access allowed for revoked role`)
	}
}
