package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//create a new type of 'Deck' which is slice of strings
type deck []string

//func [(receiver)] [name of the function]
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

//Creating a new function
//func [nameOfTheFunction]() [return_type]
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
