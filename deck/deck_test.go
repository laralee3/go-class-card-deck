package deck

// func TestCardCreationSuccess(t *testing.T) {
// 	newCard := Card{Suit: "spade", Rank: "two"}

// 	if newCard.Suit != "spade" || newCard.Rank != "two" {
// 		t.Errorf("failed to create a Deck with expected suit and rank")
// 	}

// 	newCardName := newCard.CardName()
// 	if newCardName != "two of spades" {
// 		t.Errorf("CardName() failed to return expected name")
// 	}
// }

// func TestEmptyDeckFail(t *testing.T) {
//     newDeck := &Deck{}

//     err := newDeck.validate()
//     if err == nil {
//         t.Error("deck validation failed to detect duplicate card")
//     }
// }
