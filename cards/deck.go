package main

import "fmt"

// Create a new type of 'deck which is a slice of strings
type deck []string

// Create a new print function to loop through deck of cards, and print out the value of each card
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}