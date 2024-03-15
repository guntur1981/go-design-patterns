package main

import (
	"fmt"
)

type Permission struct {
	canCreate, canRead, canUpdate, canDelete bool
}

type User struct {
	Name       string
	Permission *Permission
}

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

	liz := bob.DeepCopy()
	liz.Name = "Liz"
	liz.Permission.canDelete = true

	fmt.Println(bob.Permission)
	fmt.Println(liz.Permission)
}
