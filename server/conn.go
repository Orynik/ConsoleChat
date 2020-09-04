package server

import "github.com/Orynik/ConsoleChat/cli"

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

//CLIConn Класс - клиент
type CLIConn struct{}

//NewCLIConn Конструктор
func NewCLIConn(opts *CLIConnOps) *cli.Client {
	return &cli.Client{}
}

//Send  Отправка сообщения
func (c *CLIConn) Send(_ []byte) {}
