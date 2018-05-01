package deck

import (
	"fmt"
	"math/rand"
	"strings"
)

// Card is composed of a suit and a rank
type Card struct {
	suit string
	rank string
}

func (c Card) cardName() string {
	return c.rank + " of " + c.suit
}

// Deck are slices of Card structs (undealt and dealt)
type Deck struct {
	cards []Card
	dealt []Card
}

// CardsLeft checks and prints number of cards left undealt
func (d *Deck) CardsLeft() {
	fmt.Printf("Cards Left: %d", len(d.cards))
}

// Deal deals a card from the beginning of the deck and moves it to dealt
func (d *Deck) Deal() {
	fmt.Println(strings.Title(d.cards[0].cardName()))
	d.dealt = append(d.dealt, d.cards[0])
	d.cards = d.cards[1:]
}

// PrintDeck prints the undealt cards
func (d *Deck) PrintDeck() {
	for _, x := range d.cards {
		fmt.Println(strings.Title(x.cardName()))
	}
}

// ShuffleDeck merges the dealt cards back in (if necessary) and shuffles
func (d *Deck) ShuffleDeck() {
	d.cards = append(d.cards, d.dealt...)

	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// GenerateDeck generates a deck with given ranks and suits []string
func GenerateDeck(ranks, suits []string) Deck {
	newDeck := Deck{}

	for _, x := range suits {
		for _, y := range ranks {
			newDeck.cards = append(newDeck.cards, Card{x, y})
		}
	}

	return newDeck
}
