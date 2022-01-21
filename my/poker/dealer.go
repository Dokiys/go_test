package poker

import (
	"github.com/pkg/errors"
)

type Dealer struct {
	stack *stack
}

func NewDealer() *Dealer {
	dealer := &Dealer{}
	dealer.CutTheDeck()
	return dealer
}

func (self *Dealer) CutTheDeck() {
	self.stack = &stack{
		dc: deckCards,
		v:  RandAlphabetStr(10),
	}
}

func (self *Dealer) Deal() (*HandCard, error) {
	stack := self.stack

	stack.l.Lock()
	defer stack.l.Unlock()

	if stack.c >= 17 {
		return nil, errors.New("超出发牌数量限制！")
	}

	i := 0
	cards := [3]int{}
	for k, v := range stack.dc {
		{
			if v == "" {
				continue
			}
			cards[i] = k
			stack.dc[k] = ""
		}
		if i++; i >= 3 {
			break
		}
	}

	stack.c++
	return &HandCard{cards: cards, v: stack.v}, nil
}

func (self *Dealer) Compare(c1 *HandCard, c2 *HandCard) (bool, error) {
	if c1.v != c2.v {
		return false, errors.New("比较的不是同一副牌！")
	}
	return c1.score() > c2.score(), nil
}
