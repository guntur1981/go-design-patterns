package main

import "fmt"

type Beverage interface {
	Cost() int
	Description() string
}

type Milk struct {
	cost int
}

func (m Milk) Cost() int {
	return m.cost
}

func (m Milk) Description() string {
	return "milk"
}

type Tea struct {
	cost int
}

func (t Tea) Cost() int {
	return t.cost
}

func (t Tea) Description() string {
	return "tea"
}

// decorator
type MilkTea struct {
	milk *Milk
	tea  *Tea
}

func (mt MilkTea) Cost() int {
	return mt.milk.Cost() + mt.tea.Cost()
}

func (mt MilkTea) Description() string {
	return "milk with tea"
}

func ChooseBeverage(b Beverage) {
	fmt.Printf("You choose %s for $%d.\n", b.Description(), b.Cost())
}

func main() {
	m := Milk{3}
	t := Tea{2}
	mt := MilkTea{&m, &t}
	ChooseBeverage(mt)
}
