package cmd

import (
	"fmt"
	scanner "gomap/src/scanner"
	"time"
)

type TcpScanCommand struct {
	CIDR string `short:"c" long:"cidr" description:"Target CIDR to scan" required:"true"`
	Port string `short:"p" long:"port" description:"Ports ranges to scan" required:"true"`
}

func (tsc *TcpScanCommand) Execute(args []string) error {

	s := scanner.New("tcp", tsc.CIDR, tsc.Port, time.Microsecond * 100000)
	s.Scan()

	fmt.Println(s)

	return nil
}