package std

import (
	"hash/fnv"
	"testing"
)

func hash(s string) int64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int64(h.Sum32())
}

func TestStrHash(t *testing.T) {
	t.Log(hash("123"))
	t.Log(hash("123"))
}
