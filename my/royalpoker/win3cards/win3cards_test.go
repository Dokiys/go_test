package win3cards

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestWin3Cards(t *testing.T) {
	poker := NewPoker()
	poker.CutTheDeck()
	handCard, _ := poker.Deal()
	t.Log(handCard)
}

func BenchmarkWin3Cards(b *testing.B) {
	poker := NewPoker()
	m := map[string]int{"Leopard": 0, "RoyalFlush": 0, "Flush": 0, "Straight": 0, "Pair": 0, "Single": 0}
	for i := 0; i < b.N; i++ {
		poker.CutTheDeck()
		handCard, err := poker.Deal()
		assert.NoError(b, err)
		if handCard.isLeopard() {
			m["Leopard"]++
		} else if handCard.isRoyalFlush() {
			m["RoyalFlush"]++
		} else if handCard.isFlush() {
			m["Flush"]++
		} else if handCard.isStraight() {
			m["Straight"]++
		} else if handCard.isPair() {
			m["Pair"]++
		} else {
			m["Single"]++
		}
	}

	b.Log()
	b.Logf("测试次数：%d==============",b.N)
	seq := []string{"Leopard", "RoyalFlush", "Flush", "Straight", "Pair", "Single"}
	var data [][]string
	for _, t := range seq {
		r := fmt.Sprintf("%f", float64(m[t]) / float64(b.N) * 100)
		rate := r[:5] + "%"
		data = append(data, []string{t, strconv.Itoa(m[t]), rate})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type","Times", "Rate"})
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
