package main

import "fmt"

type Memento interface {
	GetBalance() int
}

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

type MementoBalance struct {
	balance int
}

func (m MementoBalance) GetBalance() int {
	return m.balance
}

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
