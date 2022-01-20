package poker

import "github.com/pkg/errors"

type Dealer struct {
	s *stack
}

func (self *Dealer) CutTheDeck() {
	self.s = &stack{
		dc: deckCards,
		v:  RandAlphabetStr(10),
	}
}

func (self *Dealer) Deal() (*HandCard, error) {
	s := self.s

	s.l.Lock()
	defer s.l.Unlock()

	if s.c >= 17 {
		return nil, errors.New("超出发牌数量限制！")
	}

	i := 0
	cards := [3]int{}
	for k, v := range s.dc {
		{
			if v == "" {
				continue
			}
			cards[i] = k
			s.dc[k] = ""
		}
		if i++; i >= 3 {
			break
		}
	}

	s.c++
	return &HandCard{cards: cards, v: s.v}, nil
}

func (self *Dealer) Compare(c1 *HandCard, c2 *HandCard) (bool, error) {
	if c1.v != c2.v {
		return false, errors.New("比较的不是同一副牌！")
	}
	return c1.score() > c2.score(), nil
}
