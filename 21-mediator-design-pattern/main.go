package main

import "fmt"

type Room struct {
	users []*User
}

func (r *Room) Join(name string) *User {
	user := &User{name, r}
	r.users = append(r.users, user)
	return user
}

func (r *Room) Broadcast(from *User, msg string) {
	for _, user := range r.users {
		if user != from {
			fmt.Printf("[%s] to [%s]: %s\n", from.GetName(), user.GetName(), msg)

		}
	}
}

func (r *Room) SendMessage(from, to *User, msg string) {
	fmt.Printf("[%s] to [%s]: %s\n", from.GetName(), to.GetName(), msg)
}

type User struct {
	name string
	room *Room
}

func (u User) GetName() string {
	return u.name
}

func main() {
	room := &Room{}

	alice := room.Join("Alice")
	bob := room.Join("Bob")
	room.Join("Simon")

	room.SendMessage(alice, bob, "Hi Bob!")
	room.Broadcast(bob, "Hello All!")
}
