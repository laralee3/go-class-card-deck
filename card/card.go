package card

import "errors"

// Card is composed of a suit and a rank
type Card struct {
	rank string
	suit string
}

// CardName returns the name of the card
func (c *Card) CardName() string {
	return c.rank + " of " + c.suit
}

// Validate card struct
func (c *Card) validate() error {
	if c.rank == "" {
		return errors.New("rank must not be empty")
	}
	if c.suit == "" {
		return errors.New("suit must not be empty")
	}
	return nil
}

// GenerateCard returns a new card
func GenerateCard(rank, suit string) Card {
	return Card{rank, suit}
}
