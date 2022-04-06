package std

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"
)

type Ap struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
	D int32 `json:"d"`
	E int32 `json:"e,omitempty"`
	F F     `json:"f,omitempty"`
}
type F struct {
	Fa int32 `json:"fa,omitempty"`
	Fb int32 `json:"fb,omitempty"`
}

// TestJsonOmitEmptyMarshal 测试初始值不生成对应json结构
// Go中结构体转Json可以通过`omitempty`来设置零值的字段，不转成Json。
// 但是需要注意的是：本身就为零值的值也会被忽略(如果一个成员变量没有被赋值，本身也就是其对应的零值)；
// 并且空的结构体同样会被转成对应的字段，即使其所有成员变量都被添加了`omitempty`，依旧会以空结构的形式出现在Json字符串中
func TestJsonOmitEmptyMarshal(t *testing.T) {
	ap := Ap{
		A: 1, B: 2, D: 3,
		E: 0,
		F: F{},
	}
	bytes, _ := json.Marshal(ap)
	t.Log(string(bytes)) // {"a":1,"b":2,"d":3,"f":{}}
}

// TestJsomMarshalMap map转json
func TestJsomMarshalMap(t *testing.T) {
	m := make(map[int32]string, 10)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	j, _ := json.Marshal(m)
	t.Log(string(j)) // {"1":"a","2":"b","3":"c"}
}

// TestJsonUnmarshalMap 解析map json
func TestJsonUnmarshalMap(t *testing.T) {
	//str := `{'1':'a','2':'b','3':'c'}` 不支持
	str := `{"1":"a","2":"b","3":"c"}`
	var m map[int32]string

	_ = json.Unmarshal([]byte(str), &m)
	t.Log(m) // map[1:a 2:b 3:c]
}

// TestJsonUnmarshalInterface 解析num Interface
func TestJsonUnmarshalInterface(t *testing.T) {
	str := `{"a":"a","b":2}`
	type S struct {
		A interface{} `json:"a"`
		B interface{} `json:"b"`
	}
	s := &S{}

	_ = json.Unmarshal([]byte(str), &s)
	t.Logf("%T", s.A) // string
	t.Logf("%T", s.B) // float64

	t.Log(int64(s.B.(float64)))
}

type Config struct {
	Type         string `json:"type"`
	Value        string `json:"value"`
	Form         string `json:"form"`
	DefaultValue string `json:"default_value"`
}

// TestJsonOverCopy 测试json多出到字段解析到struct
func TestJsonOverCopy(t *testing.T) {
	conf := Config{}
	s := "{\"axis\": [\"B2\"],\"type\": \"type_a\", \"value\": \"123\",\"form\": \"%Y/%m/%d\"}"
	json.Unmarshal([]byte(s), &conf)

	t.Log(conf)
}

func JsonCopy(source interface{}, target interface{}) error {
	bytes, _ := json.Marshal(source)

	err := json.Unmarshal(bytes, target)
	return errors.Wrap(err, "数据拷贝失败！")
}

// TestJsonCopyStruct
func TestJsonCopyStruct(t *testing.T) {
	type A struct {
		Id   int
		Name string
	}

	var target A
	source := &A{
		Id:   1,
		Name: "zhangsan",
	}

	err := JsonCopy(source, &target)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("target: %v", target)
}

// TestJsonSlice json序列化切片
func TestJsonSlice(t *testing.T) {
	s1 := []string{}        // []
	s2 := make([]string, 2) // ["",""]
	var s3 []string         // null
	b1, _ := json.Marshal(s1)
	b2, _ := json.Marshal(s2)
	b3, _ := json.Marshal(s3)

	t.Log(string(b1))
	t.Log(string(b2))
	t.Log(string(b3))
}
