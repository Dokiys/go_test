package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var value = flag.String("v", "[]", "传入Json格式的SendConfig数组")
var field = flag.String("f", "all", "选择需要输出的字段")

// pbpaste | go run ./myparse.go
func main() {
	flag.Parse()
	if *value != "[]" {
		err := DoParse(*value, *field)
		if err != nil {
		    fmt.Println(err.Error())
		}
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		*value = scanner.Text()
		err := DoParse(*value, *field)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
