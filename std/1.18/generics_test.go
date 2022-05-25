package __18

import (
	"fmt"
	"golang.org/x/example/stringutil"
	"testing"
)

func Add[T int | int32](a T, b T) T {
	return a + b
}
func TestGenerics(t *testing.T) {
	t.Log(Add(1,2))
}

// go test -fuzz=Fuzz
func FuzzGenerics(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, b int) {
		//if Add(a,b) == 20 {
		if Add[int](a,b) == 20 {
			t.Error("Error 20")
		}
	})
}

func TestWorkspace(t *testing.T) {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("Hello"))
}