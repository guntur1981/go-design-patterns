# Description

The Memento Design Pattern is a behavioral design pattern that allows you to capture and save an object's internal state so that it can be restored later.

## Example

Consider the following example:

```
package main

import "fmt"

type BankAccount struct {
	balance int
	history []int
}

func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{balance, make([]int, 0)}
}

func (b *BankAccount) Deposit(amount int) {
	b.history = append(b.history, b.balance)
	b.balance += amount
}

func (b BankAccount) GetBalance() int {
	return b.balance
}

func (b *BankAccount) Undo() {
	if len(b.history) == 0 {
		return
	}
	b.balance = b.history[len(b.history)-1]
	b.history = b.history[:len(b.history)-1]
}

func main() {
	b := NewBankAccount(100)
	fmt.Printf("Balance 1 = %d\n", b.GetBalance())

	b.Deposit(50)
	fmt.Printf("Balance 2 = %d\n", b.GetBalance())

	b.Deposit(25)
	fmt.Printf("Balance 3 = %d\n", b.GetBalance())

	// undo last deposit
	b.Undo()
	fmt.Printf("Balance 4 = %d\n", b.GetBalance())
}
```

In the example above, we simulate saving the state (the balance) of `BankAccount` every time a deposit is made.

## The Reason Why This Is Not a Good Practice

1. **Breaks Single Responsibility Principle**: The `BankAccount` not only responsible for managing the balance, but also responsible for undoing logic.
2. **Breaks Open-Closed Principle**: Extending the functionality (e.g., adding redo functionality) would further complicate the logic.
3. **Tight Coupling**: The state saving logic is tightly coupled with the deposit functionality. It makes memory consumption higher due to the decision to save state does not take place externally (deterimined by the user).

## A Better Approach

**First**, let's create an interface `Memento` to be used by `BankAccount` along with its implementation:

```
type Memento interface {
	GetBalance() int
}

type MementoBalance struct {
	balance int
}

func (m MementoBalance) GetBalance() int {
	return m.balance
}
```

**Second**, remove all tight-coupled undoing logic from `BankAccount` and make it only restore from a `Memento`:

```
type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{balance}
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
}

func (b BankAccount) GetBalance() int {
	return b.balance
}

func (b *BankAccount) Restore(m Memento) {
	if m != nil {
		b.balance = m.GetBalance()
	}
}
```

**Third**, for undoing (and future redoing) logic, create a `MementoList`:

```
type MementoList struct {
	history []Memento
	curPos  int
}

func NewMementoList() *MementoList {
	return &MementoList{make([]Memento, 0), -1}
}

func (ml *MementoList) Store(balance int) {
	ml.curPos++
	ml.history = append(ml.history[:ml.curPos], &MementoBalance{balance})
}

func (ml *MementoList) Undo() Memento {
	if ml.curPos < 0 {
		return nil
	}
	m := ml.history[ml.curPos]
	ml.curPos--
	return m
}
```

**Finally**, we can use the Memento Design Pattern like this:

```
func main() {
	ml := NewMementoList()
	b := NewBankAccount(100)
	fmt.Printf("Balance 1 = %d\n", b.GetBalance())

	b.Deposit(50)
	ml.Store(b.GetBalance()) // intentionally store the balance
	fmt.Printf("Balance 2 = %d\n", b.GetBalance())

	b.Deposit(25)
	fmt.Printf("Balance 3 = %d\n", b.GetBalance())

	// undo last deposit
	b.Restore(ml.Undo())
	fmt.Printf("Balance 4 = %d\n", b.GetBalance())
}
```
