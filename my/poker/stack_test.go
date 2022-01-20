package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	dealer := Dealer{}
	dealer.CutTheDeck()
	cards, err := dealer.Deal()
	assert.NoError(t, err)
	t.Log(cards)
	t.Log(deckCards)
}
