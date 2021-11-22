package stringer

import "testing"

func TestStringer(t *testing.T) {
	t.Logf("One string: %s", One.String())
}
