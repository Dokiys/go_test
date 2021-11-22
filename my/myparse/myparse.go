package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

const (
	myparseTypUp = "up"
	myparseTypLow = "low"
)

var value = flag.String("v", "", "传入需要转换的字符串")
var tpy = flag.String("t", "up", "选择需要输出的类型:up|low")

// go build ./myparse.go
func main() {
	flag.Parse()
	if *value != "" {
		//err := erparams.DoParse(*value, *tpy)
		err := doParse(*value, *tpy)
		if err != nil {
		    fmt.Println(err.Error())
		}
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		*value = scanner.Text()
		err := doParse(*value, *tpy)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func doParse(str string, typ string) error {
	switch typ {
	case myparseTypUp:
		fmt.Println(strings.ToTitle(str))
	case myparseTypLow:
		fmt.Println(strings.ToLower(str))
	default:
		return errors.New("Unsupported type!")
	}
	return nil
}
