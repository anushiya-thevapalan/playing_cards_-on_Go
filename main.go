package main

func main() {
	//var cards string = "Ace of Spades"
	//Equivalent to the above method of declaration
	//cards := newCard()

	//Arrays are of fixed length. Slices are capable of varying the length
	//Declaring slice
	//cards := []string{"Ace of Hearts", newCard()}

	// cards := newDeck()
	// cards.saveToFile("mycards")
	cards := newDeckFromFile("mycards")

	hand, remainingCards := deal(cards, 6)
	hand.print()
	remainingCards.print()
}
