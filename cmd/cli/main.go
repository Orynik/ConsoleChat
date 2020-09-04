package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/Orynik/ConsoleChat/cli"
)

var (
	CliOpts *cli.CLI = new(cli.CLI)
)

func init() {
	flag.StringVar(&CliOpts.IP, "ip", "127.0.0.1", "IP address tcp web-server")
	flag.StringVar(&CliOpts.Port, "p", "8080", "web port tcp web-server's")
}

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", CliOpts.IP+":"+CliOpts.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")
}
