![build passing](https://travis-ci.com/erichnascimento/nanosec.svg?branch=master)
[![codecov](https://codecov.io/gh/erichnascimento/nanosec/branch/master/graph/badge.svg)](https://codecov.io/gh/erichnascimento/nanosec)
# Introduction
In progress

```go
package main

import (
	"fmt"

	"github.com/erichnascimento/nanosec/pkg/resource"
	"github.com/erichnascimento/nanosec/storage"
)

func main() {
  // Resources
  s := storage.NewKeyValueStorage()
  resources := resource.NewResourceStorage(s)
  resources.AddResource("org.corporation.add")
  resources.RenameResource(r1.Name, "org.corporation.create")
  resources.GetResource(r2.Name)
  resources.DeleteResource(r3.Name)

  // Roles
  s = storage.NewKeyValueStorage()
  roles := role.NewRoleStorage(s)
  roles.AddRole("root")
  roles.RenameRole("root", "admin")
  roles.GetRole("admin")
  roles.DeleteRole("admin")

  // Permissions
  s = storage.NewKeyValueStorage()
  permissions := permission.NewPermissionStorage(s)

  permissions.Grant("org.corporation.add", ["admin"])
  permissions.CheckPermission("org.corporation.add", ["admin"])
  permissions.Revoke("org.corporation.add", ["admin"])

  // Credentials
  s = storage.NewKeyValueStorage()
  credentials := credential.NewCredentialStorage(s)

	userScheme := credentials.NewUserScheme("username")
	cm := credentials.NewCredentialManager(userScheme)
	cm.SetPassword("password")
	cm.AddRoles([]string{"admin", "manager"})
	cm.RemoveRoles([]string{"root"})

  credentials.AddCredentialForUser("erichnascimento", "1234", ["admin"])
  credentials.ChangeUserPassword("erichnascimento", "old", "new")
  credentials.AddUserRole("erichnascimento", ["hr", "finance"])
  credentials.RemoveUserRole("erichnascimento", ["admin"])

  s = storage.NewKeyValueStorage()
  sessions := session.NewSessionStorage(s)
  token := sessions.AuthorizeUser("username", "password")
  session.Authenticate(token)
  session.Destroy(token)

  entityName, entityID := org.Corporation.GetEntity()
  granted := permission.TokenHasAccess(token, "org.corporation.add")
  activity.Append("erichnascimento", "org.corporation.add", entityName, entityID, AccessGrantedAsComment(granted))
}
```

```
Output:
Added resource: org.corporation.add
Renamed resource from: org.corporation.add to org.corporation.create
Getting resource: org.corporation.create
Resource: org.corporation.create deleted sucessfully
```
