package main

import (
	Server "calculator/src"
	"fmt"
)

func main() {

	server_settings := Server.Server{
		Port: 8080,
		Host: "localhost",
	}

	fmt.Println(server_settings.Port)

	router := Server.Router()

	server_settings.RunServer(router)
}
