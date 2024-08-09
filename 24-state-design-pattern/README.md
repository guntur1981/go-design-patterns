# Description

The State Design Pattern is a behavioral design pattern that allows an object to change its behavior when its internal state changes. This pattern is particularly useful when an object must change its behavior at runtime depending on its state, and it helps in managing state transitions cleanly and systematically.

## Example

Consider the following example:

```
package main

import "fmt"

type State int

const (
	Unplugged State = iota
	Plugged
	Charging
)

type Action int

const (
	Unplug Action = iota
	Plug
	StartCharging
	StopCharging
)

type EVCharger struct {
	state State
}

func NewEVCharger() *EVCharger {
	return &EVCharger{Unplugged}
}

func (c *EVCharger) Do(a Action) {
	switch a {
	case Unplug:
		switch c.state {
		case Unplugged:
			fmt.Println("It's already unplugged!")
		case Plugged:
			fmt.Println("Charger has been unplugged.")
			c.state = Unplugged
		case Charging:
			fmt.Println("Please stop charging before unplugged the charger!")
		}

	case Plug:
		switch c.state {
		case Unplugged:
			fmt.Println("Charger has been plugged.")
			c.state = Plugged
		case Plugged:
			fmt.Println("It's already plugged!")
		case Charging:
			fmt.Println("Charging stopped due to the charger was plugged back in.")
			c.state = Plugged
		}

	case StartCharging:
		switch c.state {
		case Unplugged:
			fmt.Println("Please plug the charger before start charging.")
		case Plugged:
			fmt.Println("Charging...")
			c.state = Charging
		case Charging:
			fmt.Println("It's already charging!")
		}

	case StopCharging:
		switch c.state {
		case Unplugged:
			fmt.Println("Charger is unplugged!")
		case Plugged:
			fmt.Println("Not charging currently!")
		case Charging:
			fmt.Println("Charging has stopped.")
			c.state = Plugged
		}
	}
}

func main() {
	charger := NewEVCharger()
	charger.Do(Plug)
	charger.Do(Unplug)
	charger.Do(StartCharging)
}
```

In the above example, we simulate an EV (Electric Vehicle) charger which has multiple states and its states change depending on the given action.

## The Reason Why This Is Not a Good Practice

1. **Breaks Single Responsibility Principle**: The `EVCharger` has two responsibilities: managing states and handling state-specific behaviour.
2. **Breaks Open-Closed Principle**: If another state is added, The `Do()` method must be modified to handle the new state.

## A Better Approach

**First**, let's change `Action` constant to an interface:

```
type Action interface {
	Plug()
	Unplug()
	StartCharging()
	StopCharging()
}
```

Here, we will enforce the implementation of all registered actions for each state. For example, here is the concrete implementation for `Unplugged` state:

```
type UnpluggedAction struct {
	charger *EVCharger
}

func (u *UnpluggedAction) Plug() {
	fmt.Println("Ready to charge!")
	u.charger.setState(Plugged)
}

func (u *UnpluggedAction) Unplug() {
	fmt.Println("It's already unplugged!")
}

func (u *UnpluggedAction) StartCharging() {
	fmt.Println("Please plug the charger before start charging.")
}

func (u *UnpluggedAction) StopCharging() {
	fmt.Println("Charger is unplugged!")
}
```

Note that, each action now stores the `EVCharger` object. So, when the action is appropriate, the `EVCharger`'s state will be changed accordingly.

**Second**, modify the `EVCharger` object to allow storing actions:

```
type EVCharger struct {
	actions      map[State]Action
	currentState State
}
```

And, here's how we initialise it:

```
func NewEVCharger() *EVCharger {
	charger := &EVCharger{}
	charger.actions = map[State]Action{
		Unplugged: &UnpluggedAction{charger: charger},
		Plugged:   &PluggedAction{charger: charger},
		Charging:  &ChargingAction{charger: charger},
	}

	charger.currentState = Unplugged

	return charger
}
```

**Third**, to adhere Open-Closed Principle, remove `Do()` method and replace it with separated methods for all actions:

```
func (c *EVCharger) Plug() {
	c.actions[c.currentState].Plug()
}

func (c *EVCharger) Unplug() {
	c.actions[c.currentState].Unplug()
}

func (c *EVCharger) StartCharging() {
	c.actions[c.currentState].StartCharging()
}

func (c *EVCharger) StopCharging() {
	c.actions[c.currentState].StopCharging()
}

```

**Finally**, we can use the State Design Pattern like this:

```
func main() {
	charger := NewEVCharger()

	charger.Plug()
	charger.Unplug()
	charger.StartCharging()
}
```
