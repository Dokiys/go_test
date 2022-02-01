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
		dc: map[int]string{
			2: "S2", 3: "S3", 4: "S4", 5: "S5", 6: "S6", 7: "S7", 8: "S8", 9: "S9", 10: "S10", 11: "SJ", 12: "SQ", 13: "SK", 14: "SA",
			102: "H2", 103: "H3", 104: "H4", 105: "H5", 106: "H6", 107: "H7", 108: "H8", 109: "H9", 110: "H10", 111: "HJ", 112: "HQ", 113: "HK", 114: "HA",
			202: "C2", 203: "C3", 204: "C4", 205: "C5", 206: "C6", 207: "C7", 208: "C8", 209: "C9", 210: "C10", 211: "CJ", 212: "CQ", 213: "CK", 214: "CA",
			302: "D2", 303: "D3", 304: "D4", 305: "D5", 306: "D6", 307: "D7", 308: "D8", 309: "D9", 310: "D10", 311: "DJ", 312: "DQ", 313: "DK", 314: "DA",
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
	return HandCard{Cards: cards, v: stack.v}, nil
}

