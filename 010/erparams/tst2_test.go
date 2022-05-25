package main

import (
	"github.com/pkg/errors"
	"regexp"
	"testing"
)

func Conv(str string) (int, error) {
	regp, err := regexp.Compile("[0-9]*")
	if err != nil {
		return 0, errors.Wrapf(err, "参数非法")
	}
	m := regp.Find([]byte(str))
	if len(m) == 0 || len(m)!= len(str) {
		return 0, errors.New("参数非法")
	}

	var result int
	for _, r := range string(m) {
		result = result*10 + (int(r)-48)
	}

	return result, nil
}

func TestConv(t *testing.T) {
	cases := []struct {
		name string
		ins string
		exp int
	}{
		{
			"纯数字测试",
			"123",
			123,
		},
	}
	for _, c := range cases {
		r, err := Conv(c.ins)
		if err != nil {
			t.Errorf("测试用例%s结果出错:%v",c.name, err)
		}
		if r != c.exp {
			t.Errorf("测试用例%s结果出错，期望：%d实际：%d",c.name,c.exp,r)
		}
	}
}
