package main

import (
	"fmt"

	"github.com/ajay-patidar-0/rag/internal/rag"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("RAG based project")
	godotenv.Load()
	query := "what is java"

	_, err := rag.GetEmbedding(query)
	if err != nil {
		fmt.Printf("error at GetEmbedding %v", err)
	}
}
