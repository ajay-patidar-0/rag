package main

import (
	"fmt"

	"github.com/ajay-patidar-0/rag/internal/api"
)

func main() {
	fmt.Println("RAG based project")
	server := api.NewApiServer()

	server.Run()
}
