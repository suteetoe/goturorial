package main

import (
	"fmt"

	"github.com/suteetoe/goturorial/gopackage/book"
)

func main() {

	b := book.NewBook()
	fmt.Printf("%T, %v\n", b, b)

	b.Name = "Ring"
	fmt.Printf("%T, %v\n", b, b)
}
