package server

import (
	"fmt"
	"net"

	"github.com/Orynik/ConsoleChat/cli"
)

//CLIConnOps Структура для хранения параметров клиента
type CLIConnOps struct {
	// Функция обработки сообщения
	onMessage func()
	// Функция, которая вызывается при ошибке отправки,
	// чтения сообщений и т.д
	onError func()
	/*{
	// сообщить серверу, что меня нужно удалить
	}*/
}

var mapRequest map[string]string = map[string]string{
	"reg": "needreg",
}

//CLIConn Класс - клиент
type CLIConn struct {
	client *cli.Client
}

//NewCLIConn Конструктор
func NewCLIConn(opts *CLIConnOps) *cli.Client {
	return &cli.Client{}
}

//Send  Отправка сообщения
func (c *CLIConn) Send(_ []byte) {}

//Read  Отправка сообщения
func (c *CLIConn) Read(conn net.Conn, clientData CLIConn) {
	defer conn.Close()
	for {
		//TODO: Решить проблему с длинной строки считывания
		input := make([]byte, 1024*4)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Print("Read error: ", err)
			break
		}
		source := string(input[0:n])

		if clientData.client.Login == "" {
			fmt.Printf("%s try unauthorization send message. \n", conn.RemoteAddr().String())
			conn.Write([]byte("Нужно зарегистрироваться или авторизоваться перед тем, как отправлять сообщения.\nЧтобы авторизоваться введите /a, чтобы зарегистрироваться - /reg"))
		}

		fmt.Println(source)

		conn.Write([]byte(source))
	}
}
