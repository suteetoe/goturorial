package main

import "fmt"

func main() {
	InterfaceImprementAssertion()
}

func InterfaceTypeAssertion() {

	// interface type assert tion

	var a interface{}

	a = 10

	fmt.Printf("%T, %v\n", a, a)

	// var i int
	// i = a // error but [] cannot use a (type interface {}) as type int in assignment: need type assertion]
	// i = a.(float32) // cannot use a.(float32) (type float32) as type int in assignment

	if i, ok := a.(float32); ok {
		fmt.Printf("%T, %v\n", i, i) // ถ้า type ไม่ตรงก็ไม่เข้า
	}

	if i, ok := a.(int); ok {
		fmt.Printf("%T, %v\n", i, i)
	}

}

func InterfaceImprementAssertion() {
	var a, b Obj

	a = rectangle{w: 4, l: 4}
	b = triangle{h: 4, w: 4}

	fmt.Println(a.Area())

	fmt.Println(b.Area())

	if v, ok := b.(triangle); ok {
		// assert b is a triangle
		v.h = 5
		fmt.Println(v.Area())
	}

	fmt.Println(a.Area())
	fmt.Println(b.Area())
}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.w * rec.w
}

type triangle struct {
	w, h float64
}

func (tri triangle) Area() float64 {
	return tri.h * tri.w * 0.5
}
