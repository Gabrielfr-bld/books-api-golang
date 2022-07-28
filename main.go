package main

import (
	"github.com/Gabrielfr-bld/books-api-golang/server"
	"github.com/Gabrielfr-bld/books-api-golang/database"
)

func main() {
	database.StartDB()

	s := server.NewServer()

	s.Run()
}