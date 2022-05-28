package main

func main() {
	cards := newDeck()
	// cards.saveToFile("my_cards")
	// deck := newDeckFromFile("my_cards")
	// deck.print()

	// rand.Seed(time.Now().UnixNano())

	// shuffleDeck(cards).print()

	cards.shuffle()
	cards.print()
}
