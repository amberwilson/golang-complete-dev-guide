package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	suits := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	ranks := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck, error) {
	if handSize > len(d) {
		return nil, nil, errors.New(fmt.Sprintf("Cannot deal %v cards, there are only %v in the deck.", handSize, len(d)))
	}

	return d[:handSize], d[handSize:], nil
}

func (d deck) toString() string {
	return strings.Join(d, "|")
}

func deckFromString(str string) deck {
	return deck(strings.Split(str, "|"))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deckFromString(string(content))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// region MY SHUFFLE
func shuffleDeck(d deck) deck {
	shuffled := deck{}
	for len(d) > 1 {
		// get a random card that's left in the deck
		cardIndex := randomCardIndex(d)

		// add the card to the shuffled set
		shuffled = append(shuffled, d[cardIndex])

		// slice out selected card
		d = append(d[:cardIndex], d[cardIndex+1:]...)
	}

	// add last card to shuffled set
	shuffled = append(shuffled, d[0])

	return shuffled
}

func randomCardIndex(d deck) int {
	min := 0
	max := len(d) - 1

	return rand.Intn(max-min) + min
}

// endregion MY SHUFFLE
