package main

import (
	"github.com/rightjoin/aqua"
	cubereum "sblock/api/service"
)

// main function
func main() {
	server := aqua.NewRestServer()
	server.AddService(&cubereum.Sblock{})
	server.Run()
}
