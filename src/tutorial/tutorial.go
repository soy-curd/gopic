package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	a := []byte{0, 255}
	b := []byte{0x61, 0x62}
	// a := []byte{256}
	// b := []byte{-1}
	fmt.Printf("%d, %d\n", a, b)
	fmt.Printf("%x, %x\n", a, b)
	fmt.Printf("%q, %q\n", a, b)
	print(hex.Dump([]byte("helloこんにちは")))
	c := []byte{0, 1, 2, 3}
	d := []byte{0, 1, 2, 3}
	// print(c == d)            // false
	println(bytes.Equal(c, d)) // true
	e := bytes.NewBuffer([]byte("helloこんにちは"))
	next, _ := e.ReadByte()
	fmt.Printf("%q, %d\n", next, e.Len())
	next, _ = e.ReadByte()
	fmt.Printf("%q, %d\n", next, e.Len())
	next, _ = e.ReadByte()
	fmt.Printf("%q, %d\n", next, e.Len())
	// "h", 19
	// "e", 18
	// "l", 17

}
