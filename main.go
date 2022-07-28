package main

import (
	"github.com/Gabrielfr-bld/books-api-golang/server"
)

func main() {
	s := server.NewServer()

	s.Run()
}