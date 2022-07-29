package tcli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
)

const VERSION = "v1.0.0"

var notifySignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

func RegisterNotifySignal(signals ...os.Signal) {
	notifySignals = signals
}

func Run(cleanup func(os.Signal), cmds ...*cli.Command) {
	var stop = make(chan struct{})

	app := cli.NewApp()
	app.Name = "tcli"
	app.Usage = "tcli hello"
	app.UsageText = "tcli [global options] command [command options] [arguments...]"
	app.Copyright = "Copyright ©2022 Dokiy"
	app.Version = VERSION
	app.Authors = []*cli.Author{{
		Name:  "Dokiy",
		Email: "dokiychang@gmail.com",
	}}
	app.Commands = cmds
	app.After = func(c *cli.Context) error {
		close(stop)
		return nil
	}

	go waitSign(stop, cleanup)
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
	<-stop
}

func waitSign(stop chan<- struct{}, cleanup func(os.Signal)) {
	n := make(chan os.Signal)
	signal.Notify(n, notifySignals...)

	cleanup(<-n)
	stop <- struct{}{}
}
