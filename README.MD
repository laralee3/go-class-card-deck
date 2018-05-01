Card/Deck generator in go for go-class

## Command reference
```
go run
go build
go test ./...
//... = test all subpackages recursively
```


## Andrew's aliases
```
alias goct='go test -cover -timeout 30s'
alias gocvt='go test -cover -v -timeout 30s'
alias gocrt='go test -cover -race -timeout 30s'
alias gocrvt='go test -cover -race -v -timeout 30s'
```

## Andrew's examples
```
func main() {
    log.SetPrefix("{go-cards} ")

    rand.Seed(time.Now().UTC().UnixNano())

    deck := cards.NewDeck()
    printDeckOfCards(deck)

    log.Print("*********************\n")
    log.Print("Shuffling the deck...\n")
    log.Print("*********************\n")
    deck.Shuffle()
    printDeckOfCards(deck)
}
```

```
func main() {
    log.SetPrefix("{go-cards} ")

    rand.Seed(time.Now().UTC().UnixNano())

    deck := cards.NewDeck()
    deck.Print()

    log.Print("*********************\n")
    log.Print("Shuffling the deck...\n")
    log.Print("*********************\n")

    deck.Shuffle()
    deck.Print()
}
```

```
func TestDeckCreationSucceed(t *testing.T) {
    newDeck := NewDeck()

    if newDeck.GetCount() != NumberOfCards {
        t.Errorf("failed to create a Deck with %d cards, got %d cards instead", NumberOfCards, newDeck.GetCount())
    }

    err := newDeck.validate()
    if err != nil {
        t.Errorf("deck validation failed: %v", err)
    }
}

func TestEmptyDeckFail(t *testing.T) {
    newDeck := &Deck{}

    err := newDeck.validate()
    if err == nil {
        t.Error("deck validation failed to detect duplicate card")
    }
}
```