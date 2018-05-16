package card

import (
	"testing"
)

func TestCardStructCreationSuccess(t *testing.T) {
	newCardStruct := Card{rank: "two", suit: "spades"}

	if newCardStruct.suit != "spades" || newCardStruct.rank != "two" {
		t.Error("failed to create a Card with expected suit and rank")
	}

	err := newCardStruct.validate()
	if err != nil {
		t.Errorf("Card validation failed: %v", err)
	}
}

func TestCardCreationSuccess(t *testing.T) {
	newCardGenerated := GenerateCard("two", "spades")

	if newCardGenerated.suit != "spades" || newCardGenerated.rank != "two" {
		t.Error("failed to create a Card with expected suit and rank")
	}

	if newCardGenerated.CardName() != "two of spades" {
		t.Error("CardName() failed to return expected name")
	}

	err := newCardGenerated.validate()
	if err != nil {
		t.Errorf("Card validation failed: %v", err)
	}
}

func TestCardValidateErrorSuccess(t *testing.T) {
	newCardNoRank := Card{suit: "spades"}

	errNoRank := newCardNoRank.validate()
	if errNoRank == nil {
		t.Error("Card validation failed to find missing rank")
	}

	newCardNoSuit := Card{rank: "two"}

	errNoSuit := newCardNoSuit.validate()
	if errNoSuit == nil {
		t.Error("Card validation failed to find missing suit")
	}
}

func TestCardCreationFail(t *testing.T) {
	newCard := GenerateCard("two", "spades")

	err := newCard.validate()
	if err == nil {
		t.Error("Card validation failed to find problems")
	}
}
