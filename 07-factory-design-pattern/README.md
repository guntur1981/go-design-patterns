# Description

The Factory Design Pattern allows you to create objects without exposing the creation logic to the user.

The difference with The Builder Design Pattern is that the Builder creates objects **step by step**, while the Factory creates them **in one go**.

## Example

Consider the following example of user permissions for a blog post:

```
package main

import "fmt"

const (
	Create int = iota
	Read
	Update
	Delete
)

type Permission struct {
	CanCreate, CanRead, CanUpdate, CanDelete bool
}

func (p *Permission) IsAllow(operation int) bool {
	allowed := false
	switch operation {
	case Create:
		allowed = p.CanCreate
	case Read:
		allowed = p.CanRead
	case Update:
		allowed = p.CanUpdate
	case Delete:
		allowed = p.CanDelete
	}

	return allowed
}

func main() {
	// Bob as an admin
	bob := Permission{CanCreate: false, CanRead: true, CanUpdate: false, CanDelete: true}

	// Liz as an author
	liz := Permission{CanCreate: true, CanRead: true, CanUpdate: true, CanDelete: false}

	// Tom as a guest
	tom := Permission{CanCreate: false, CanRead: true, CanUpdate: false, CanDelete: false}

	fmt.Println("Is Bob allowed to create a post:", bob.IsAllow(Create))
	fmt.Println("Is Liz allowed to create a post:", liz.IsAllow(Create))
	fmt.Println("Is Tom allowed to create a post:", tom.IsAllow(Create))
}
```

In the above example, we created Bob, Liz, and Tom permissions by setting their permission properties at initialization.

## The Reason Why This Is Not a Good Practice

1. If the Permission struct is in different package, we should adhere to the Dependency Inversion Principle where the high-level module should not access low-level module directly except through interfaces.
2. If Permission initialization is repeated elsewhere in the code, we might set incorrect permission accidentally.
3. The read permission from above example is always set to true, so we want to have default read permission without having to set it at every initialization.

## A Better Approach

Let's create a factory to help initialize the Permission struct based on the roles.

**First**, let's add some Role options:

```
const (
	Admin int = iota
	User
	Guest
)
```

**Second**, let's add an interface to only expose the IsAllow() behavior and hide the properties. Notice the difference between `Permission` and `permission`.

```
type Permission interface {
	IsAllow(operation int) bool
}

type permission struct {
	canCreate, canRead, canUpdate, canDelete bool
}

func (p *permission) IsAllow(operation int) bool {
	allowed := false
	switch operation {
	case Create:
		allowed = p.canCreate
	case Read:
		allowed = p.canRead
	case Update:
		allowed = p.canUpdate
	case Delete:
		allowed = p.canDelete
	}

	return allowed
}
```

**Finally**, let's create a Permission factory based on roles.

```
func NewPermission(role int) Permission {
	// default zero values are false
	p := permission{}

	// set it initially as a guest
	// where the default value for the read permission is always true
	p.canRead = true

	switch role {
	case Admin:
		p.canDelete = true
	case User:
		p.canCreate = true
		p.canRead = true
		p.canUpdate = true
	}
	return &p
}
```
