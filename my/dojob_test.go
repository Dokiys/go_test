package main

import (
	"flag"
	"reflect"
	"testing"
)

type Worker struct {}

func (self Worker) HelloWork() bool {
	print("Hello Work!\n")
	return true
}

func TestDojob(t *testing.T) {
	var mStr string
	flag.StringVar(&mStr, "m", "", "执行方法名")

	flag.Set("m", "HelloWork")
	flag.Parse()
	workerValue := reflect.ValueOf(&Worker{})

	m := workerValue.MethodByName(mStr)
	if !m.IsValid() {
		panic("Method Not Found!")
	}

	value := m.Call([]reflect.Value{})
	ok := value[0].Bool()
	if ok {
		print("OK!")
	}
}
