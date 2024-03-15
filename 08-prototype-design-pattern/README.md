# Description

The Prototype Design Pattern allows you to create objects by copying from an existing object. In other words, **object cloning** or **deep copying**.

## Example

Consider the following example:

```
package main

import "fmt"

type Permission struct {
	canCreate, canRead, canUpdate, canDelete bool
}

type User struct {
	Name       string
	Permission Permission
}

func main() {
	bob := User{
		Name: "Bob",
		Permission: Permission{
			canCreate: false,
			canRead:   true,
			canUpdate: false,
			canDelete: false,
		},
	}

	liz := bob // shallow copy
	liz.Name = "Liz"
	liz.Permission.canDelete = true

	fmt.Println(bob)
	fmt.Println(liz)
}
```

In the above example, we created object Liz by (shallow) copying from object Bob and then making necessary changes according to Liz's requirements.

## The Reason Why This Is Not a Good Practice

If the Permission property is changed to a pointer, the Liz's permisions cannot be changed without affecting Bob's permission.

```
package main

import "fmt"

type Permission struct {
	canCreate, canRead, canUpdate, canDelete bool
}

type User struct {
	Name       string
	Permission *Permission // changed to a pointer
}

func main() {
	bob := User{
		Name: "Bob",
		Permission: &Permission{
			canCreate: false,
			canRead:   true,
			canUpdate: false,
			canDelete: false,
		},
	}

	liz := bob // shallow copy
	liz.Name = "Liz"
	liz.Permission.canDelete = true

	fmt.Println(bob.Permission)
	fmt.Println(liz.Permission)
}
```

## A Better Approach

Let's create a deep copy method for the User object:

```
func (u User) DeepCopy() *User {
	user := User{
		Name: u.Name,
		Permission: &Permission{
			canCreate: u.Permission.canCreate,
			canRead:   u.Permission.canRead,
			canUpdate: u.Permission.canUpdate,
			canDelete: u.Permission.canDelete,
		},
	}
	return &user
}
```

And we can use the DeepCopy() method as follows:

```
	liz := bob.DeepCopy()
	liz.Name = "Liz"
	liz.Permission.canDelete = true

	fmt.Println(bob.Permission)
	fmt.Println(liz.Permission)
```
