package win3cards

import (
	"github.com/dokiy/royalpoker/common"
	"github.com/pkg/errors"
)

type Win3Cards struct {
	stack *stack
}

func NewPoker() *Win3Cards {
	w3c := &Win3Cards{}
	w3c.CutTheDeck()
	return w3c
}

func (self *Win3Cards) CutTheDeck() {
	self.stack = &stack{
		dc: deckCards,
		v:  common.RandAlphabetStr(10),
	}
}

func (self *Win3Cards) Deal() (HandCard, error) {
	stack := self.stack

	stack.l.Lock()
	defer stack.l.Unlock()

	if stack.c >= 17 {
		return HandCard{}, errors.New("超出发牌数量限制！")
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
	return HandCard{Cards: cards, V: stack.v}, nil
}

