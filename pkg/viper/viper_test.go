package viper

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperWrite(t *testing.T) {
	viper.SetConfigType("yaml")
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	viper.WriteConfigAs("config")
}

func TestViperRead(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.ReadInConfig()
	assert.Equal(t, 35, viper.Get("age"))
}

func TestViperUnmarshal(t *testing.T) {
	vp := viper.New()
	vp.SetConfigType("yaml")
	err := vp.ReadConfig(strings.NewReader(`
age: 18
name: zhangsan
`))
	assert.NoError(t, err)
	type People struct {
		Age  int    `yaml:"age"`
		Name string `yaml:"name"`
	}
	p := new(People)
	err = vp.Unmarshal(p, func(conf *mapstructure.DecoderConfig) {
		conf.TagName = "yaml"
	})
	assert.NoError(t, err)

	assert.Equal(t, 18, p.Age)
}
