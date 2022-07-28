package main

import (
	"fmt"
	"testing"
)

type Client interface {
	Request()
}

type Destroyer interface {
	Destroy()
}

var _ Client = (*Impl)(nil)
var _ Destroyer = (*Impl)(nil)

type Impl struct{}

type Config struct {
	a int
	b int32
	c int64
	d string
}

type ClientFactory func(id int) Client

//func BuildClientImpl(id int) Client {
func BuildClientImpl(id int) *Impl {
	conf := doSomeGetConf(id)
	return NewImpl(conf)
}

//type A struct {
//	BuildMethod func(id int) *Impl
//}
//
//func TestBuildClientImpl(t *testing.T) {
//	//a := &A{BuildClientImpl(1)}
//	a := &A{BuildMethod: BuildClientImpl}
//	a.BuildMethod(1).Request()
//}

type A struct {
	cf ClientFactory
}

func TestBuildClientImpl(t *testing.T) {
	a := &A{cf: NewClientFactory()}
	a.cf(1).Request()
}

func NewClientFactory() ClientFactory {
	return func(id int) Client {
		conf := doSomeGetConf(id)
		return NewImpl(conf)
	}
}

//func NewImpl(a int, b int32, c int64, d string) *Impl {
//func NewImpl(conf Config) Client {
func NewImpl(conf Config) *Impl {
	fmt.Printf("%d,%d,%d,%s", conf.a, conf.b, conf.c, conf.d)
	return &Impl{}
}

type Option func(impl *Impl)

func SetA(a int) Option {
	return func(impl *Impl) {
		fmt.Printf("Set %d\n", a)
	}
}

func NewImpl2(opts ...Option) *Impl {
	impl := &Impl{}
	for _, opt := range opts {
		opt(impl)
	}
	return impl
}

func TestImpl2(t *testing.T) {
	impl2 := NewImpl2(SetA(1))
	impl2.Request()
}

//func NewClient(id int) *Impl {
//	conf := doSomeGetConf(id)
//	fmt.Printf("%d,%d,%d,%s", conf.a, conf.b, conf.c, conf.d)
//	return &Impl{}
//}

//func NewDestroyer() Destroyer {
func NewDestroyer() *Impl {
	return &Impl{}
}

func (c *Impl) Request() {
	fmt.Println("Impl Request")
}

func (c *Impl) Destroy() {
	panic("Impl Destroy")
}

/*
	配置参数
*/
type MConfig map[int]Config

func (m MConfig) init() {
	m[1] = Config{a: 1, b: 1, c: 1, d: "1"}
	m[2] = Config{a: 2, b: 2, c: 2, d: "2"}
}

var m MConfig

func TestFactory(t *testing.T) {
	m = make(MConfig)
	m.init()
	conf1 := doSomeGetConf(1)
	//var conf = Config{a: 1,b: 1,c: 1,d: "1"}
	//c1 := NewClientC(conf)
	c1 := NewImpl(conf1)
	c1.Request()

	conf2 := doSomeGetConf(2)
	c2 := NewImpl(conf2)
	c2.Request()

	/*
		Factory Create Client
	*/
	BuildClientImpl(1).Request()
	BuildClientImpl(1).Destroy()
}

// 获取配置参数
func doSomeGetConf(id int) Config {
	return m[id]
}
