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
  s := storage.NewKeyValueStorage()
  resources := resource.NewResourceStorage(s)

  r1, _ := resources.AddResource("org.corporation.add")
  fmt.Println("Added resource:", r1.Name)

  r2, _ := resources.RenameResource(r1.Name, "org.corporation.create")
  fmt.Println("Renamed resource from:", r1.Name, "to", r2.Name)

  r3, _ := resources.GetResource(r2.Name)
  fmt.Println("Getting resource:", r3.Name)

  err := resources.DeleteResource(r3.Name)
  if err == nil {
    fmt.Println("Resource:", r3.Name, "deleted sucessfully")
  }

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
