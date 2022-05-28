package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	size := 52

	// a deck is created with 52 cards
	if len(d) != size {
		t.Errorf("Expected deck length of %v, but got %v", size, len(d))
	}

	// first card is an Ace of Clubs
	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be \"Ace of Clubs\", but got %v", d[0])
	}

	// last card is a King of Spades
	if d[size-1] != "King of Spades" {
		t.Errorf("Expected first card to be \"King of Spades\", but got %v", d[size-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of %v, but got %v", 52, len(deck))
	}

	os.Remove(filename)
}

// region MY TESTS
func TestDealReturnsAnErrorIfAskingForTooManyCards(t *testing.T) {
	deck := newDeck()

	_, _, err := deal(deck, 100)

	if err == nil {
		t.Errorf("Expected an error because requesting 100 cards when deck only has 52, but no error received")
	}
}

// endregion MY TESTS
