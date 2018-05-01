package main

import (
	"log"

	"github.com/glympse/go-class/deck"
)

func main() {
	log.SetPrefix("{go-class cards} ")

	ranks := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}
	suits := []string{"spades", "hearts", "clubs", "diamonds"}

	newDeckOfCards := deck.GenerateDeck(ranks, suits)

	log.Print("\n******* Deck of Cards *******")
	newDeckOfCards.PrintDeck()

	log.Print("\n\n******* Randomized Deck *******")
	newDeckOfCards.ShuffleDeck()
	newDeckOfCards.PrintDeck()

	log.Print("\n\n******* Deal card *******")
	newDeckOfCards.Deal()
	newDeckOfCards.CardsLeft()

	log.Print("\n\n******* Deal another card *******")
	newDeckOfCards.Deal()
	newDeckOfCards.CardsLeft()

	log.Print("\n\n******* Check remaining deck (cheater) *******")
	newDeckOfCards.PrintDeck()

	log.Print("\n\n******* Reshuffle Deck *******")
	newDeckOfCards.ShuffleDeck()
	newDeckOfCards.PrintDeck()
}
