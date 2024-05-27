# Description

The Mediator Design Pattern is a component that facilitates communication between other components without them necessarily being aware of each other or having direct (reference) access to each other

## Example

Consider the following example:

```
package main

import "fmt"

type Room struct {
	users []*User
}

func (r *Room) Join(user *User) {
	r.users = append(r.users, user)
}

type User struct {
	name string
	room *Room
}

func NewUser(name string, room *Room) *User {
	return &User{name, nil}
}

func (u User) GetName() string {
	return u.name
}

func (u User) SendMessage(to *User, msg string) {
	fmt.Printf("[%s] to [%s]: %s\n", u.name, to.GetName(), msg)
}

func main() {
	room := &Room{}
	alice := NewUser("Alice", room)
	bob := NewUser("Bob", room)

	room.Join(alice)
	room.Join(bob)

	alice.SendMessage(bob, "Hi Bob!")
	bob.SendMessage(alice, "Hello Alice!")
}
```

In the above example, we have a chat roow where users can send message to each other.

## The Reason Why This Is Not a Good Practice

1. **Violates Single Responsibility Principle**: Each object has multiple responsibilities, including communication logic.
2. **Increased Complexity**: As the number of interaction grows, the system becomes more complex and harder to manage. For example, if the user want to broadcast a message, then the user must loop through all users available to send the message.
3. **Tight Coupling**: Direct references between objects increase coupling, making the system harder to maintain and extend.

## A Better Approach

**First**, let's make the struct `Room` be the mediator:

```
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
```

**Second**, we move the communication responsibilities from `User` to the mediator `Room`:

```
type User struct {
	name string
	room *Room
}

func (u User) GetName() string {
	return u.name
}
```

**Finally**, we can implement the Mediator Design Pattern like this:

```
func main() {
	room := &Room{}

	alice := room.Join("Alice")
	bob := room.Join("Bob")
	room.Join("Simon")

	room.SendMessage(alice, bob, "Hi Bob!")
	room.Broadcast(bob, "Hello All!")
}
```
