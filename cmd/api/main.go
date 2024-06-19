package main

import (
	"fmt"

	"github.com/EspenS93/filterview/internal/server"
)

func main() {
	server := server.New()
	server.RegisterRoutes()
	err := server.Run()

	fmt.Println(err)
}
