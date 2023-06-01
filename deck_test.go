package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()
	assert.Len(t, deck1.cards, 52)
	assert.Len(t, deck2.cards, 52)

	same := true
	for i := range deck1.cards {
		same = same && (deck1.cards[i] == deck2.cards[i])
	}
	assert.False(t, same)
}

func TestDraw(t *testing.T) {
	deck := NewDeck()

	cards := deck.Draw(5)
	assert.Len(t, cards, 5)
	assert.False(t, deck.Empty())

	deck.Draw(52 - 5)
	assert.True(t, deck.Empty())
}

func TestEmpty(t *testing.T) {
	deck := NewDeck()
	assert.False(t, deck.Empty())

	deck.Draw(51)
	assert.False(t, deck.Empty())

	deck.Draw(1)
	assert.True(t, deck.Empty())
}

func TestNewDeckNoShuffle(t *testing.T) {
	deck := NewDeckNoShuffle()
	assert.Len(t, deck.cards, 52)
	hand := deck.Draw(5)
	t.Log(hand)
}

func TestDrawWithRng(t *testing.T) {
	seed := []byte{
		0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0,
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x00,
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
	}
	deck := NewDeckNoShuffle()
	t.Log(deck.cards)
	cards := deck.DrawWithRng(5, seed)
	assert.Len(t, cards, 5)
	t.Log(cards)
	t.Log(RankClass(Evaluate(cards)))
}
