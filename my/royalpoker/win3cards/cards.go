package win3cards

type HandCard struct {
	Cards [3]int
	v string
}

func (self HandCard) Version() string {
	return self.v
}

func Compare(h1 HandCard,h2 HandCard) bool {
	return h1.Score() > h2.Score()
}

func (self HandCard) Score() (score int) {
	if self.isLeopard() {
		score = 500
	} else if self.isRoyalFlush() {
		score = 400
	} else if self.isFlush() {
		score = 300
	} else if self.isStraight() {
		score = 200
	} else if self.isPair() {
		score = 100
	} else {
		score = 0
	}
	return (score+self.baseScore())*10
}

func (self HandCard) baseScore() int {
	x, y, z := sortCard(self.Cards[0]%100, self.Cards[1]%100, self.Cards[2]%100)
	if isStraight(x, y, z) && x == 2 {
		return x + y + z - 13
	}
	return x + y + z
}

func (self HandCard) suitScore() int {
	x, y, z := self.Cards[0], self.Cards[1], self.Cards[2]
	var baseMax int
	if x%100 >= y%100 && x%100 >= z%100 {
		baseMax = x
	}
	if y%100 >= x%100 && y%100 >= z%100 {
		baseMax = y
	}
	if x%100 >= y%100 && x%100 >= z%100 {
		baseMax = z
	}

	var suitMax int
	if x == baseMax && x/100 >= suitMax {
		suitMax = x / 100
	}
	if y == baseMax && y/100 >= suitMax {
		suitMax = y / 100
	}
	if z == baseMax && z/100 >= suitMax {
		suitMax = z / 100
	}

	return suitMax
}

func (self HandCard) isLeopard() bool {
	x, y, z := self.Cards[0], self.Cards[1], self.Cards[2]
	return isLeopard(x, y, z)
}

func (self HandCard) isRoyalFlush() bool {
	x, y, z := self.Cards[0], self.Cards[1], self.Cards[2]
	return isFlush(x, y, z) && isStraight(x, y, z)
}

func (self HandCard) isFlush() bool {
	x, y, z := self.Cards[0], self.Cards[1], self.Cards[2]
	return isFlush(x, y, z) && !isStraight(x, y, z)
}

func (self HandCard) isStraight() bool {
	x, y, z := self.Cards[0], self.Cards[1], self.Cards[2]
	return isStraight(x, y, z) && !isFlush(x, y, z)
}

func (self HandCard) isPair() bool {
	x, y, z := self.Cards[0]%100, self.Cards[1]%100, self.Cards[2]%100
	return isPair(x, y, z) && !isLeopard(x, y, z)
}

// ============================================================

func isLeopard(x, y, z int) bool {
	return x%100 == y%100 && y%100 == z%100
}

func isFlush(x, y, z int) bool {
	return x/100 == y/100 && y/100 == z/100
}

func isStraight(x, y, z int) bool {
	x, y, z = sortCard(x%100, y%100, z%100)
	return (x+1 == y && x+2 == z) || (x+1 == y && x+12 == z)
}

func isPair(x, y, z int) bool {
	return (x == y || x == z || y == z)
}

func sortCard(a, b, c int) (x, y, z int) {
	x, y, z = a, b, c
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}

	return
}
