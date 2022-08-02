package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"syscall"

	"github.com/Dokiys/gstp"
	"github.com/urfave/cli/v2"
)

const VERSION = "v1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "gsfp"
	app.Usage = "gsfp -s='tem*' template.go"
	app.UsageText = "gsfp [global options] [arguments...]"
	app.Copyright = "Copyright ©2022 Dokiy"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "s", Value: ".*", Usage: "regexp match struct"},
	}
	app.Authors = []*cli.Author{{
		Name:  "Dokiy",
		Email: "dokiychang@gmail.com",
	}}
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func action(ctx *cli.Context) error {
	s := ctx.String("s")
	exp, err := regexp.Compile(s)

	if err != nil {
		fmt.Printf("-s invalid: %s\n", err)
		return nil
	}

	for _, src := range ctx.Args().Slice() {
		f, err := os.Open(src)
		if err != nil {
			if errors.Is(err, syscall.ENOENT) {
				continue
			}

			fmt.Println(err)
			return nil
		}

		w := os.Stdout
		if err := gstp.GenProto(f, w, *exp); err != nil {
			return err
		}
	}

	return nil
}
