package main

import (
	"fmt"
	"os"

	"tcli"
)

func main() {
	signals := []os.Signal{}
	tcli.RegisterNotifySignal(signals...)

	cmd := tcli.NewHelloCmd()

	cleanup := func(signal os.Signal) {
		fmt.Println(signal)
	}

	tcli.Run(cleanup, cmd)
}
