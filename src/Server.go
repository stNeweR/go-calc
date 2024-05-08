package src

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port int
	Host string
}

func (s *Server) RunServer(router *http.ServeMux) {
	addr := fmt.Sprintf(":%d", s.Port)

	fmt.Println("Сервер запущен по адресу: ", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}
