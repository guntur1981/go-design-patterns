# Description

The Interpreter Design Pattern is a behavioral design pattern that processes structured text of data. Does so by turning it into separate lexical tokens (lexing) and then interpreting sequences of said tokens (parsing).

## Example

Consider the following example:

```
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluateExpression(s string) int {
	result := 0
	operand := strings.Builder{}
	var operator string
	tokens := []string{}

	pushToken := func(token string) bool {
		if len(token) > 0 {
			tokens = append(tokens, token)
			return true
		}
		return false
	}

	// lexing
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '+', '-':
			if pushToken(operand.String()) {
				operand.Reset()
			}
			operator = string(c)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if pushToken(operator) {
				operator = ""
			}
			operand.WriteByte(c)
		}
	}
	pushToken(operand.String())

	// parsing
	haveLeft := false
	left, right := 0, 0
	for _, token := range tokens {
		switch token {
		case "+", "-":
			operator = token
		default:
			val, _ := strconv.Atoi(token)
			if !haveLeft {
				left = val
				haveLeft = true
			} else {
				right = val
				switch operator {
				case "+":
					result = left + right
				case "-":
					result = left - right
				}
				left = result
			}
		}
	}
	return result
}

func main() {
	expr := "10 - 11 + 2"
	result := evaluateExpression(expr)

	fmt.Println(result)
}
```

The above example shows how to interpret a mathemetical expression into an integer result.

## The Reason Why This Is Not a Good Practice

1. **Violates Single Responsibility Principle**: The function `evaluateExpression()` does to much - it lexing the expression and parsing it.
1. **Violates Open-Closed Principle**: If we need to add new operations (e.g., multiplication, division) it requires modification of function.

## A Better Approach

**First**, let's adhere to Single Responsibility Principle by separating the lexing and parsing processes. For lexing requirements, we define token information, as follows:

```
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
```

Then, we use it in function `Lex()`:

```
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
```

**Second**, specify the type `Operation` to be used by the parsing function.

```
type Operation int

const (
	None Operation = iota
	Addition
	Subtraction
)
```

Later, we can easily add more operations if needed, such as multiplication and division.

**Third**, create an interface `Operand` also to be used by the parsing function:

```
type Operand interface {
	Value() int
}
```

Here we implement the interface `Operand` for `Integer` and `BinaryOperation`:

```
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
```

Later, we can also add more operands if needed.

**Fourth**, Now we can define the parsing function `Parse`:

```
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
```

**Finally**, we can use Interpreter Design Pattern like this:

```
func main() {
	expr := "10 - 11 + 2"
	tokens := Lex(expr)
	parsed := Parse(tokens)
	fmt.Println("Result of", expr, "is", parsed.Value())
}
```
