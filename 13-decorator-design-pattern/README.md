# Description

The Decorator Design Pattern is a structural design that facilitates the addition of behaviors to individual objects through embedding.

## Example

Consider the following example:

```
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

func ChooseBeverage(b Beverage) {
	fmt.Printf("You choose %s for $%d.\n", b.Description(), b.Cost())
}

func main() {
	m := Milk{2}
	t := Tea{3}
	ChooseBeverage(m)
	ChooseBeverage(t)
}
```

In the above example, we have structs `Milk` and `Tea` that implement the interface `Beverage`.

Suppose we want to have a combination beverage `MilkTea`, we might do it like this:

```
// combination
type MilkTea struct {
	Milk
	Tea
}

func (mt MilkTea) Cost() int {
	return 7
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
	mt := MilkTea{m, t}
	ChooseBeverage(mt)
}
```

## The Reason Why This Is Not a Good Practice

1. We only embed the behaviours of `Milk` and `Tea` without their data `cost`. So, initializing `Milk` and `Tea` is useless.
2. It's hard to maintain consistency since it defines its own cost.

## A Better Approach

It's better to embed existing objects `Milk` and `Tea` as fields in `MilkTea`:

```
// decorator
type MilkTea struct {
	milk *Milk
	tea  *Tea
}
```

Next, modify the method `Cost()` to use the existing `Milk` and `Tea` costs:

```
func (mt MilkTea) Cost() int {
	return mt.milk.Cost() + mt.tea.Cost()
}
```

Now, we can initialize the decorator `MilkTea` like this:

```
func main() {
	m := Milk{3}
	t := Tea{2}
	mt := MilkTea{&m, &t}
	ChooseBeverage(mt)
}
```

So whenever we want to add beverages and combine them, we can create new decorators following this pattern.
