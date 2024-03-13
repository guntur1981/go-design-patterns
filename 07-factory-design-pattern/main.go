package main

import "fmt"

const (
	Admin int = iota
	User
	Guest
)

const (
	Create int = iota
	Read
	Update
	Delete
)

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

func main() {
	// Bob as an admin
	bob := NewPermission(Admin)

	// Liz as an author
	liz := NewPermission(User)

	// Tom as a guest
	tom := NewPermission(Guest)

	fmt.Println("Is Bob allowed to create a post:", bob.IsAllow(Create))
	fmt.Println("Is Liz allowed to create a post:", liz.IsAllow(Create))
	fmt.Println("Is Tom allowed to create a post:", tom.IsAllow(Create))
}
