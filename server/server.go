package server

import (
	"context"
	"fmt"
	"net"

	"github.com/Orynik/ConsoleChat/cli"
)

//ServerOpts Структура для хранения параметров сервера
type ServerOpts struct {
	ServerPort        string
	DatabasePort      string
	DatabaseName      string
	DatabaseTableName string
	DatabaseUser      string
	DatabasePass      string
}

//Server Класс - сервер
type Server struct {
	Users []cli.Client
}

//NewServer Конструктор
func NewServer() *Server {
	return &Server{}
}

//Run Запуск сервера
func (s *Server) Run(ctx context.Context, opts *ServerOpts) {

	listener, err := net.Listen("tcp", ":"+opts.ServerPort)

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Server serve on :" + opts.ServerPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte("Welcome to ConsoleChat on Golang web-server!"))
	}

}
