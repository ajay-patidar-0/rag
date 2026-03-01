package api

import (
	"log"
	"net/http"

	"github.com/ajay-patidar-0/rag/internal/db"
	"github.com/ajay-patidar-0/rag/internal/store"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ApiServer struct {
	Db *store.VectorDB
}

func NewApiServer() *ApiServer {
	db, err := db.NewDb()
	if err != nil {
		log.Fatalf("Error creating database connection : %v ", err)
	}
	return &ApiServer{
		Db: store.NewVectorDB(db),
	}
}

func (as *ApiServer) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/query", as.QueryHandler).Methods("POST", "OPTIONS")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":4000", handler))
}

func ServHome(w http.ResponseWriter, r *http.Request) {

}
