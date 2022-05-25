package alg

import (
	"math/rand"
	"testing"
)

func TestFind(t *testing.T) {
	slice := make([]int,0,10000)
	for i:=0;i < 10000;i++{
		if i == 555{
			slice = append(slice,11111)
		}else{
			slice = append(slice,rand.Intn(10000))
		}
	}

	if i, ok := Find(slice,11111,4); ok{
		t.Logf("Find at: %v\n", i)
	}else {
		t.Fatalf("Should return true, but got false!")
	}
}

func TestNotFind(t *testing.T) {
	slice := make([]int,0,10000)
	for i:=0;i < 10000;i++{
		slice = append(slice,rand.Intn(10000))
	}

	if _, ok := Find(slice,11111,4); !ok{
		t.Logf("Not find target")
	}else {
		t.Fatalf("Should not find target,but found")
	}
}