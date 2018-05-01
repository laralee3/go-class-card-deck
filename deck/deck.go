package deck

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/glympse/go-class/card"
)

// Deck are slices of Card structs (undealt and dealt)
type Deck struct {
	cards []card.Card
	dealt []card.Card
}

// CardsLeft checks and prints number of cards left undealt
func (d *Deck) CardsLeft() {
	log.Printf("Cards Left: %d", len(d.cards))
}

// Deal deals a card from the beginning of the deck and moves it to dealt
func (d *Deck) Deal() {
	log.Println(strings.Title(d.cards[0].CardName()))
	d.dealt = append(d.dealt, d.cards[0])
	d.cards = d.cards[1:]
}

// PrintDeck prints the undealt cards
func (d *Deck) PrintDeck() {
	for _, x := range d.cards {
		log.Println(strings.Title(x.CardName()))
	}
}

// ShuffleDeck merges the dealt cards back in (if necessary) and shuffles
func (d *Deck) ShuffleDeck() {
	rand.Seed(time.Now().UTC().UnixNano())

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
			newDeck.cards = append(newDeck.cards, card.GenerateCard(y, x))
		}
	}

	return newDeck
}
