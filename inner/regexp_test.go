package inner

import (
	"regexp"
	"testing"
)

// TestRegexp 正则使用
func TestRegexp(t *testing.T) {
	r := regexp.MustCompile("[0-9]+")
	t.Log(r.MatchString("abc123")) //true
	// 这个方法返回匹配的子串
	t.Log(r.FindString("abc123a")) //Hello World!
}

// TestRegexpOr 正则或条件
func TestRegexpOr(t *testing.T) {
	r := regexp.MustCompile("[0-9]|\\(abc\\)")

	t.Log(r.MatchString("(abc)")) //true
	t.Log(r.FindString("abc123a")) //123
}

// TestRegexpOperator 校验四则运算符
func TestRegexpOperator(t *testing.T) {
	r := regexp.MustCompile("[\\+\\-\\*/]")
	t.Log(r.FindString("a+b"))
}

// TestRegexpNum 校验数字
func TestRegexpNum(t *testing.T) {
	r := regexp.MustCompile("[0-9]")
	t.Log(r.MatchString("a+b"))
}