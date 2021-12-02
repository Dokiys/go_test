package std

import "testing"

func one() int {
	return 1
}
func f() (i int) {
	return one()
}

func TestBase(t *testing.T) {
	t.Log(f())
}
