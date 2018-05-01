package main

import (
	"fmt"

	"github.com/glympse/go-class/deck"
)

func main() {
	ranks := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}
	suits := []string{"spades", "hearts", "clubs", "diamonds"}

	newDeckOfCards := deck.GenerateDeck(ranks, suits)

	fmt.Println("\n******* Deck of Cards *******")
	newDeckOfCards.PrintDeck()

	fmt.Println("\n\n******* Randomized Deck *******")
	newDeckOfCards.ShuffleDeck()
	newDeckOfCards.PrintDeck()

	fmt.Println("\n\n******* Deal card *******")
	newDeckOfCards.Deal()
	newDeckOfCards.CardsLeft()

	fmt.Println("\n\n******* Deal another card *******")
	newDeckOfCards.Deal()
	newDeckOfCards.CardsLeft()

	fmt.Println("\n\n******* Check remaining deck (cheater) *******")
	newDeckOfCards.PrintDeck()

	fmt.Println("\n\n******* Reshuffle Deck *******")
	newDeckOfCards.ShuffleDeck()
	newDeckOfCards.PrintDeck()
}
