package rag

import (
	"database/sql"
	"fmt"

	"github.com/ajay-patidar-0/rag/internal/store"
)

func SearchSimilar(db *sql.DB, query string) ([]string, error) {
	emb, err := GetEmbedding(query)
	if err != nil {
		return nil, fmt.Errorf("error occurs at get embedding in search %v", err)
	}

	vecQuery := store.ToPGVector(emb)

	rows, err := db.Query(`SELECT content FROM questions ORDER BY embedding <=>$1 LIMIT 2`, vecQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string

	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return nil, err
		}
		result = append(result, content)
	}
	return result, nil
}
