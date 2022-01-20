package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDealerCompare(t *testing.T) {
	d := &Dealer{}
	cases := []struct {
		hc1    *HandCard
		hc2    *HandCard
		except bool
	}{
		{
			hc1:    &HandCard{cards: [3]int{14, 2, 3}},
			hc2:    &HandCard{cards: [3]int{2, 4, 3}},
			except: false,
		},
		{
			hc1:    &HandCard{cards: [3]int{14, 2, 3}},
			hc2:    &HandCard{cards: [3]int{2, 114, 3}},
			except: true,
		},
	}

	for i, c := range cases {
		result, err := d.Compare(c.hc1, c.hc2)
		assert.NoError(t, err)
		assert.Equal(t, c.except, result, i+1)
	}
}
