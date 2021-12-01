package main

import "fmt"

func main() {
	fn := newCounterFunc()

	fn2 := fn

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	fmt.Println(fn2())
}

func newCounterFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
