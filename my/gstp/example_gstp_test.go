package gstp

import (
	"bytes"
	"os"
	"regexp"
)

func ExampleGenProto() {
	src := `
package example

import (
	"fmt"
	"time"
)

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
`
	r := bytes.NewReader([]byte(src))

	exp, _ := regexp.Compile(".*")
	_ = GenProto(r, os.Stdout, *exp)
	// Output:
	//
	//// Item Comment 1
	///*
	//	Item Comment 1
	//*/
	//// Item Comment 1
	//message Item {
	//	// Item ItemId Comment 1
	//	int32 item_id = 1;
	//
	//	string name = 2;
	//
	//	Duration duration = 3;
	//
	//	google.protobuf.Timestamp created_at = 4;
	//}
	//
	//message TemplateData {
	//
	//	repeated string arr = 1;
	//
	//	repeated Item items = 2;
	//
	//	map<string,Item> map1 = 3;
	//}
}
