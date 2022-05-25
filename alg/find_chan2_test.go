package alg

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
)

func TestSearch(t *testing.T) {
	slice := make([]int,0,10000)
	for i:=0;i < 10000;i++{
		if i == 555{
			slice = append(slice,11111)
		}else{
			slice = append(slice,rand.Intn(10000))
		}
	}

	if _,err := Search2(slice,11111); err == nil{
		t.Logf("Get!")
	}else {
		t.Fatal(err)
	}
}

func Test1(t *testing.T){
	ctx,cancel := context.WithCancel(context.Background())


	go func(ctx context.Context) {
		ctx = context.WithValue(ctx, 1,"a")
		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Value(1))
	}
}

func Test2(t *testing.T) {
	type keyctxa string
	type keyctxb string
	ctx := context.Background()
	var k1 keyctxa = "KeyA"
	ctxA := context.WithValue(ctx, k1, 1)
	var k2 keyctxb = "KeyA"
	ctxB := context.WithValue(ctxA, k2, 2)

	fmt.Println(ctxA.Value(k1))
	fmt.Println(ctxB.Value(k2))
}