package std

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceEqual(t *testing.T) {
	var i interface{}
	i = "-"
	assert.Equal(t, "-", i)
	i = 0
	assert.Equal(t, 0, i)
}