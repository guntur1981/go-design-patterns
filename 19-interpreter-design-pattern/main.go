package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
)

type Token struct {
	Type TokenType
	Text string
}

func Lex(s string) []Token {
	var result []Token
	operand := strings.Builder{}
	pushOperand := func() bool {
		if operand.Len() > 0 {
			result = append(result, Token{Int, operand.String()})
			operand.Reset()
			return true
		}
		return false
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+':
			result = append(result, Token{Plus, "+"})
		case '-':
			result = append(result, Token{Minus, "-"})
		default:
			j := i // j scans for operand
			for ; j < len(s); j++ {
				if unicode.IsDigit(rune(s[j])) {
					operand.WriteByte(s[j])
				} else if pushOperand() {
					break
				}
			}
			i = j // skip what j has scanned
		}
	}
	// last operand
	pushOperand()

	return result
}

type Operation int

const (
	None Operation = iota
	Addition
	Subtraction
)

type Operand interface {
	Value() int
}

type Integer struct {
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value}
}

func (i Integer) Value() int {
	return i.value
}

// The result of a binary operation can also be an operand
type BinaryOperation struct {
	Type        Operation
	Left, Right Operand
}

func (b BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Subtraction:
		return b.Left.Value() - b.Right.Value()
	}
	return 0
}

func Parse(tokens []Token) Operand {
	result := BinaryOperation{}

	// move existing operation to the left side
	moveToLeft := func() {
		if result.Left != nil && result.Right != nil {
			i := NewInteger(result.Value())
			result.Left = i
			result.Type = None
			result.Right = nil
		}
	}

	for _, token := range tokens {
		switch token.Type {
		case Plus:
			moveToLeft()
			result.Type = Addition
		case Minus:
			moveToLeft()
			result.Type = Subtraction
		case Int:
			val, _ := strconv.Atoi(token.Text)
			i := NewInteger(val)
			if result.Left == nil {
				result.Left = i
			} else {
				result.Right = i
			}
		}
	}

	return result
}

func main() {
	expr := "10 - 11 + 2"
	tokens := Lex(expr)
	parsed := Parse(tokens)
	fmt.Println("Result of", expr, "is", parsed.Value())
}
