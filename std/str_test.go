package std

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestStrCompare 字符串比较
func TestStrCompare(t *testing.T) {
	str1 := "121"
	str2 := "122"

	t.Log(str1 > str2)
}

// TestStrToByte 测试保存字符串的interface能不能直接强转成byte
func TestStrToByte(t *testing.T) {
	var str = "123"
	var i interface{}

	i = str
	var b []byte
	b = []byte(i.(string))
	t.Log(b)
}

// TestStrReplace 字符串替换
func TestStrReplace(t *testing.T) {
	s := "dfjalskfdls"

	s = strings.ReplaceAll(s, "s", "S")
	t.Log(s)
}

// TestStrReplaceAll 字符串替换all
func TestStrReplaceAll(t *testing.T) {
	str := "jd12kld3afs234djfkal"
	prefix := make([]rune, len(str))

	index := 0
	for _, r := range str {
		if r >= '0' && r <= '9' {
			continue
		}
		prefix[index] = r
		index++
	}

	t.Logf("替换后的字符：%s", string(prefix[:index]))
}

// TestStrGenRand 生成随机字符串,带数字
func TestStrGenRand1(t *testing.T) {
	now := time.Now().UnixNano()

	buf := make([]byte, 9)
	binary.PutVarint(buf, now)

	sum := md5.Sum(buf)
	s := fmt.Sprintf("%x", sum)[:8]
	t.Log(s)
}

const str_alphabet = "abcdefghijklmnopqrstuvwxyz"

// TestStrGenRand1 生成随机字符串，不带数字
func TestStrGenRand2(t *testing.T) {
	b := make([]byte, 8)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = str_alphabet[rand.Int63()%int64(26)]
	}
	t.Log(string(b))
}

// TestStrBatchReplaceAll 字符串替换
func TestStrBatchReplaceAll(t *testing.T) {
	v := "123252432"
	for i := 0; i < len(v); i++ {
		if v[i] == '3' {
			v = v[:i] + "" + v[i+1:]
		}
	}
	t.Log(v)
}

// TestStrSlicePrint 字符串切片打印
func TestStrSlicePrint(t *testing.T) {
	strArr := []string{"1", "2", "3"}
	t.Logf("%s", strArr)
}

func TestStrToInt(t *testing.T) {
	str := "12378392469324342"
	i, err := strconv.ParseInt(str, 10, 64)
	assert.NoError(t, err)
	t.Log(i)
}

// TestStrNestedMethod 校验表达式是否有方法嵌套
func TestStrNestedMethod(t *testing.T) {
	s := "sum(a+c) + sum(a+b)"
	var isNested bool
	var exp map[string]struct{}

	for i := 0; i+4 < len(s); i++ {
		if s[i:i+4] != "sum(" && s[i:i+4] != "avg(" {
			continue
		}
		for j, c := i+4, 1; j < len(s); j++ {
			t.Log(string(s[j]))
			if s[j] == '(' {
				c++
			} else if s[j] == ')' {
				c--
			} else if j+4 < len(s) && (s[j:j+4] == "sum(" || s[j:j+4] == "avg(") {
				isNested = true
				goto BREAK
			}

			if c == 0 {
				exp[s[i+4:j]] = struct{}{}
				i = j
				break
			}
		}
	}
BREAK:
	t.Log(isNested)
	t.Log(exp)
}

//go test -run="none" -bench="Bench*"
func BenchmarkStrCheck(b *testing.B) {
	str := "君不见，黄河之水天上来，奔流到海不复回。君不见，高堂明镜悲白发，朝如青丝幕成雪"
	for i := 0; i < b.N; i++ {
		if str != "" {
			continue
		}
	}
}
func BenchmarkStrCheck2(b *testing.B) {
	str := "君不见，黄河之水天上来，奔流到海不复回。君不见，高堂明镜悲白发，朝如青丝幕成雪"
	for i := 0; i < b.N; i++ {
		if len(str) > 0 {
			continue
		}
	}
}
