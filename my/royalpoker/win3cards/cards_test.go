package win3cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	cases := []struct {
		hc     HandCard
		except int
	}{
		{
			hc:     HandCard{Cards: [3]int{14, 114, 314}},
			except: 5420,
		},
		{
			hc:     HandCard{Cards: [3]int{14, 102, 3}},
			except: 2060,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.except, c.hc.Score(), i+1)
	}

}

func TestIsLeopard(t *testing.T) {
	cases := []struct {
		hc     HandCard
		except bool
	}{
		{
			hc:     HandCard{Cards: [3]int{14, 114, 314}},
			except: true,
		},
		{
			hc:     HandCard{Cards: [3]int{14, 102, 3}},
			except: false,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.except, c.hc.isLeopard(), i+1)
	}

}

func TestIsRoyalFlush(t *testing.T) {
	cases := []struct {
		hc     HandCard
		except bool
	}{
		{
			hc:     HandCard{Cards: [3]int{102, 114, 103}},
			except: true,
		},
		{
			hc:     HandCard{Cards: [3]int{2, 114, 3}},
			except: false,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.except, c.hc.isRoyalFlush(), i+1)
	}

}

func TestIsFlush(t *testing.T) {
	cases := []struct {
		hc     HandCard
		except bool
	}{
		{
			hc:     HandCard{Cards: [3]int{14, 2, 4}},
			except: true,
		},
		{
			hc:     HandCard{Cards: [3]int{14, 2, 3}},
			except: false,
		},
		{
			hc:     HandCard{Cards: [3]int{2, 114, 3}},
			except: false,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.except, c.hc.isFlush(), i+1)
	}

}

func TestIsStraight(t *testing.T) {
	cases := []struct {
		hc     HandCard
		except bool
	}{
		{
			hc:     HandCard{Cards: [3]int{14, 102, 103}},
			except: true,
		},
		{
			hc:     HandCard{Cards: [3]int{112, 314, 213}},
			except: true,
		},
		{
			hc:     HandCard{Cards: [3]int{14, 111, 213}},
			except: false,
		},
		{
			hc:     HandCard{Cards: [3]int{2, 14, 3}},
			except: false,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.except, c.hc.isStraight(), i+1)
	}

}

func TestIsPair(t *testing.T) {
	cases := []struct {
		hc     HandCard
		except bool
	}{
		{
			hc:     HandCard{Cards: [3]int{14, 114, 302}},
			except: true,
		},
		{
			hc:     HandCard{Cards: [3]int{14, 114, 314}},
			except: false,
		},
		{
			hc:     HandCard{Cards: [3]int{14, 102, 3}},
			except: false,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.except, c.hc.isPair(), i+1)
	}

}
