package main

import (
	"fmt"
	"log"

	"github.com/ajay-patidar-0/rag/internal/rag"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("RAG based project")
	godotenv.Load()

	// db, err := db.NewDb()
	// if err != nil {
	// log.Panic(err)
	// }

	query := `Data Communication, Storage, and Performance Optimization Methods`
	content := `Routing and Switching
		Routing and switching are essential networking processes. Switching operates within a local network and forwards data based on MAC addresses. Routing connects different networks and forwards packets using IP addresses. Routers determine the best path for data using routing algorithms and tables. Efficient routing and switching ensure fast data transmission, reduced congestion, and improved overall network performance.
		Distributed DBMS
		A Distributed DBMS stores data across multiple physical locations while appearing as a single database to users. It improves reliability, availability, and scalability. Data may be fragmented or replicated across sites. Distributed databases are commonly used in large organizations and cloud systems. They allow efficient resource sharing and fault tolerance while handling large volumes of data efficiently.`

	ans, err := rag.AskLLm(query, content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
