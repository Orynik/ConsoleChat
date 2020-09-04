package main

import (
	"context"
	"flag"

	"github.com/Orynik/ConsoleChat/server"
)

var (
	serverOpts server.ServerOpts = server.ServerOpts{}
)

func init() {
	//TODO: Добавить при парсе значений ":" к портам
	//ServersOpts
	flag.StringVar(&serverOpts.ServerPort, "p", "8080", "Web-server port")
	//DatabaseOpts
	flag.StringVar(&serverOpts.DatabasePort, "dbport", "3360", "Database port")
	flag.StringVar(&serverOpts.DatabaseName, "dbname", "test", "Database name")
	flag.StringVar(&serverOpts.DatabaseTableName, "dbtable", "test", "Database table name")
	flag.StringVar(&serverOpts.DatabaseUser, "dbuser", "root", "Database user")
	flag.StringVar(&serverOpts.DatabasePass, "dbpass", "password", "Database password")
}

func main() {
	flag.Parse()

	// Запуск сервера
	ctx := context.Background()
	srv := server.NewServer()
	srv.Run(ctx, &serverOpts)
}
