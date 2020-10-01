package main

import (
	"flag"
	"fmt"
	"net"

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

	for {
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}

		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		// получем ответ
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}

}
