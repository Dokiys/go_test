package example

import (
	"fmt"
	"time"
)

func a() {
	fmt.Println()
}

// Item Comment 1
/*
	Item Comment 1
*/
// Item Comment 1
type Item struct {
	// Item ItemId Comment 3

	// Item ItemId Comment 2
	ItemId    int // Item ItemId Comment 1
	Name      string
	Duration  time.Duration
	CreatedAt time.Time
}

type TemplateData struct {
	Arr   []string
	Items []*Item
	Map1  map[string]*Item

	// Unsupported
	//TdArr [][]string
	//Map2 map[string][]*Item
	//Map3 []map[string]*Item
	//Map4 []map[*Item]string
	//Map5 []map[string][]*Item
}

func (t *TemplateData) P() {
	fmt.Println()
}
