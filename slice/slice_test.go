package main

import "testing"

func TestWordCountc(t *testing.T) {
	given := "Apple Banana Apple Banana apple"
	want := map[string]int{
		"Apple":  2,
		"Banana": 2,
		"apple":  1,
	}

	get := wordCount(given)
	if len(get) != len(want) {
		t.Errorf("string count not match %d but %d", len(want), len(get))
	}

	if get["apple"] != want["apple"] {
		t.Errorf("count not match %d but %d", get["apple"], want["apple"])
	}
}

func TestBookString(t *testing.T) {
	give := &Book{
		Name:   "Harry Potter",
		Author: "J. K. Rowling",
	}

	want := "Harry Potter by J. K. Rowling"
	if get := give.String(); get != want {
		t.Errorf("%s : %s", want, get)
	}

	want = "Harry Potter by toe"
	give.SetName("toe")
	if get := give.String(); get != want {
		t.Errorf("%s : %s", want, get)
	}
}
