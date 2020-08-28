package server

import "context"

//ServerOpts Структура для хранения параметров сервера
type ServerOpts struct {
}

//Server Класс - сервер
type Server struct{}

//NewServer Конструктор
func NewServer() *Server {}

//Run Запуск сервера
func (s *Server) Run(ctx context.Context, opts *ServerOpts) {}
