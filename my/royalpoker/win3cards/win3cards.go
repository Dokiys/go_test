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
		dc: []int{
			2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
			102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114,
			202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214,
			302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314,
		},
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

	var remainDc = stack.dc
	cards := [3]int{}
	for i := 0; i < 3; i++ {
		for {
			n := common.RandNum(len(remainDc)-i)
			if remainDc[n] == 0 {
				remainDc = append(remainDc[:n], remainDc[n+1:]...)
				continue
			}

			cards[i] = remainDc[n]
			remainDc[n] = 0
			break
		}
	}
	stack.dc = remainDc

	stack.c++
	return HandCard{Cards: cards, v: stack.v}, nil
}

