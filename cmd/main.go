package main

import (
	"fmt"

	"github.com/ajay-patidar-0/rag/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("RAG based project")
	godotenv.Load()

	server := api.NewApiServer()
	fmt.Println("server started")

	server.Run()
}
