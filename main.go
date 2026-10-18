package main

import (
	"restapi/server"
	"restapi/todo"
)

func main() {
	list := todo.NewList()
	handlers := server.NewHttpHandlers(list)
	server := server.NewServer(handlers)

	server.LaunchServer()

	todo.StupidFunc()
}