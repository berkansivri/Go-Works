package main

func main() {
	cards := newDeck()

	cards.shuffle() 
	cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("my_cards")
	
	// newDeck := newDeckFromFile("my_cards")
	// newDeck.print()
}