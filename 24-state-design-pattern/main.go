package main

import "fmt"

type State int

const (
	Unplugged State = iota
	Plugged
	Charging
)

type Action interface {
	Plug()
	Unplug()
	StartCharging()
	StopCharging()
}

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

type PluggedAction struct {
	charger *EVCharger
}

func (p *PluggedAction) Plug() {
	fmt.Println("It's already plugged!")
}

func (p *PluggedAction) Unplug() {
	fmt.Println("Charger has been unplugged.")
	p.charger.setState(Unplugged)
}

func (p *PluggedAction) StartCharging() {
	fmt.Println("Charging...")
	p.charger.setState(Charging)
}

func (p *PluggedAction) StopCharging() {
	fmt.Println("Not charging currently!")
}

type ChargingAction struct {
	charger *EVCharger
}

func (c *ChargingAction) Plug() {
	fmt.Println("Charging stopped due to the charger was plugged back in.")
	c.charger.setState(Plugged)
}

func (c *ChargingAction) Unplug() {
	fmt.Println("Please stop charging before unplugging the charger!")
}

func (c *ChargingAction) StartCharging() {
	fmt.Println("It's already charging!")
}

func (c *ChargingAction) StopCharging() {
	fmt.Println("Charging has stopped.")
	c.charger.setState(Plugged)
}

type EVCharger struct {
	actions      map[State]Action
	currentState State
}

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

func (c *EVCharger) setState(s State) {
	c.currentState = s
}

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

func main() {
	charger := NewEVCharger()

	charger.Plug()
	charger.Unplug()
	charger.StartCharging()

	charger.Plug()
	charger.StartCharging()
	charger.StopCharging()
	charger.Unplug()
}
