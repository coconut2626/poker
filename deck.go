package poker

import (
	"encoding/binary"
	"math/rand"
)

var fullDeck *Deck

func init() {
	fullDeck = &Deck{initializeFullCards()}
}

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.Shuffle()
	return deck
}

func NewDeckNoShuffle() *Deck {
	deck := &Deck{}
	deck.cards = make([]Card, len(fullDeck.cards))
	copy(deck.cards, fullDeck.cards)
	return deck
}

func (deck *Deck) Shuffle() {
	deck.cards = make([]Card, len(fullDeck.cards))
	copy(deck.cards, fullDeck.cards)
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}

func (deck *Deck) Draw(n int) []Card {
	cards := make([]Card, n)
	copy(cards, deck.cards[:n])
	deck.cards = deck.cards[n:]
	return cards
}

func (deck *Deck) DrawWithRng(n int, seed []byte) []Card {
	cards := make([]Card, n)
	if n > (len(seed) / 4) {
		return nil
	}
	for i := 0; i < n; i++ {
		idx := binary.BigEndian.Uint32(seed[i:i+4]) % uint32(len(deck.cards))
		cards[i] = deck.cards[idx]
		deck.cards = append(deck.cards[:idx], deck.cards[idx+1:]...)
	}
	return cards
}

func (deck *Deck) Empty() bool {
	return len(deck.cards) == 0
}

func initializeFullCards() []Card {
	var cards []Card

	for _, rank := range strRanks {
		for suit := range charSuitToIntSuit {
			cards = append(cards, NewCard(string(rank)+string(suit)))
		}
	}

	return cards
}
