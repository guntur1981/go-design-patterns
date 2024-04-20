package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type Encryption interface {
	Encode() string
}

// a bridge
type Encoder interface {
	Encode(string) string
}

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

type Secret struct {
	encoder Encoder
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
