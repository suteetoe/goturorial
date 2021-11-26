package main

import (
	"fmt"
	"strings"
)

func main() {
	cde()
}

func abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed
	selectedFruit := s[:3]
	fmt.Print(s)
	// [apple banana coconut]
	fmt.Print(selectedFruit)
}

func efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed
	selectedFruit := s[4:]
	fmt.Print(selectedFruit)
	// * [elderberries figs grapes]
}

func cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed

	selectedFruit := s[2:5]
	fmt.Print(selectedFruit)
	// * [coconut durian elderberries]
}

// TODO: split words and count them
func wordCount(s string) map[string]int {

	wordCount := map[string]int{}

	words := strings.Fields(s)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount

	// split := strings.Split(s, " ")
	// result := map[string]int{}
	// for i := 0; i < len(split); i++ {
	// 	result[split[i]] = result[split[i]] + 1
	// }
	// return result
}

type Book struct {
	Name   string
	Author string
}

func (book Book) String() string {
	return fmt.Sprintf("%s by %s", book.Name, book.Author)
}

func (book *Book) SetName(name string) {
	book.Author = name
}

/*
func (book Book) String() string {

 }

 func (book *Book) SetName(name string) { }
*/
