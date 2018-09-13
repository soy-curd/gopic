package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("images/test.pgm")
	if err != nil {
		panic(err)
	}

	// err = ioutil.WriteFile("images/test.pgm", data, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("%s", hex.Dump(data[:100]))
	fmt.Println("file read/write finished!")

}
