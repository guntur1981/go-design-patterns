# Description

The Bridge Design Pattern is a mechanism that decouples an interface from an implementation. This pattern is useful for preventing the complexity explosion.

## Example

Consider the following example:

```
package main

import (
	"encoding/base64"
	"fmt"
)

type Encryption interface {
	Encode() string
}

type Secret struct {
	message string
}

func (s Secret) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(s.message))
}

func PrintEncodedMessage(e Encryption) {
	fmt.Println(e.Encode())
}

func main() {
	s := Secret{"Hello, World!"}
	PrintEncodedMessage(&s)
}
```

In the above example, we have an interface `Encryption` that has only one method `Encode()` implemented by struct `Secret` to encode its message. The encoding used is Base64.

Suppose that we want to be able to encode with hexadicemal, we might modify the code, as follows:

```
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type Encryption interface {
	EncodeBase64() string
	EncodeHex() string
}

type Secret struct {
	message string
}

func (s Secret) EncodeBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(s.message))
}

func (s Secret) EncodeHex() string {
	return hex.EncodeToString([]byte(s.message))
}
...
```

## The Reason Why This Is Not a Good Practice

By modifying the interface `Encryption` we have violated the Open-Closed Principle. And because the interface `Encryption` is used as parameter in function `PrintEncodedMessage()`, we also violate the Liskov Substitution Principle.

## A Better Approach

**First**, create a bridge interface between the interface `Encryption` and its implementation.

```
// a bridge
type Encoder interface {
	Encode(string) string
}
```

**Second**, implement all needed encodings, in this case, Base64 and Hexadecimal.

```
// implement the bride as base64
type Base64Encoder struct{}

func (b64 Base64Encoder) Encode(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

// implement the bride as hexadecimal
type Hexadecimal struct{}

func (h Hexadecimal) Encode(msg string) string {
	return hex.EncodeToString([]byte(msg))
}
```

**Finally**, use the bridge (encoding) we needed.

```
type Secret struct {
	encoder Encoder // the bridge
	message string
}

func (s Secret) Encode() string {
	return s.encoder.Encode(s.message)
}

func PrintEncodedMessage(e Encryption) {
	fmt.Println(e.Encode())
}

func main() {
	encoder := Hexadecimal{}
	s := Secret{&encoder, "Hello, World!"}
	PrintEncodedMessage(&s)
}
```
