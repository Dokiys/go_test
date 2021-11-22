//go:generate stringer -type=Num,Code
package stringer

type Num int
const (
	One   Num = iota
	Two   Num = iota
	Three Num = iota
	Four  Num = iota
	Five  Num = iota
)

type Code int32
const (
	Yi Code = 12
	Er Code = 2
)