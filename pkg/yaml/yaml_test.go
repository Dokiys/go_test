package yaml

import (
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

type Conf struct {
	I int32 `yaml:"i"`
	F float64 `yaml:"f"`
	S string `yaml:"s"`
	IsBool bool `yaml:"is_bool"`
	Arr []int32 `yaml:"arr"`
	A A `yaml:"a"`
	ManyA []A `yaml:"many_a"`
}
type A struct {
	Id int32 `yaml:"id"`
	Name string `yaml:"name"`
}

// TestYaml 测试ymal读取
func TestYaml(t *testing.T) {
	conf := Conf{}
	bytes, err := os.ReadFile("./conf.yaml")
	if err != nil {
		t.Fatal(bytes)
	}
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf)
}
