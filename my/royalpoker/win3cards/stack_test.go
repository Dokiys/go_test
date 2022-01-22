package win3cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	dealer := Win3Cards{}
	dealer.CutTheDeck()
	cards, err := dealer.Deal()
	assert.NoError(t, err)
	t.Log(cards)
	t.Log(deckCards)
}
