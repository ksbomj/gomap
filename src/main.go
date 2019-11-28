package main

import (
	"fmt"
	"github.com/ksbomj/gomap/src/cmd"
	"github.com/jessevdk/go-flags"
	"os"
)

type Opts struct {
	TcpScanCommand  cmd.TcpScanCommand  `command:"tcp"`
}

func main () {
	fmt.Println("gomap another yet network scanner")

	var opts Opts

	p := flags.NewParser(&opts, flags.Default)

	p.CommandHandler = func(command flags.Commander, args []string) error {
		err := command.Execute(args)
		if err != nil {
			fmt.Printf("[ERROR] failed with %+v", err)
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

