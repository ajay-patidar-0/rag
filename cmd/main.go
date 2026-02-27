package main

import (
	"fmt"
	"log"

	"github.com/ajay-patidar-0/rag/internal/db"
	"github.com/ajay-patidar-0/rag/internal/rag"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("RAG based project")
	godotenv.Load()
	db, err := db.NewDb()
	if err != nil {
		log.Panic(err)
	}

	query := `Data Communication, Storage, and Performance Optimization Methods`

	result, err := rag.SearchSimilar(db, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("length of result : %v ", len(result))

	for _, val := range result {
		fmt.Println(val)
	}
}
