package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

// ./cli -h
func main() {
	(&cli.App{
		Name: "hello_cli",
		Usage: "say hello to cli",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("Hello %v\n",ctx.Args().Get(0))
			return nil
		},
	}).Run(os.Args)
}


