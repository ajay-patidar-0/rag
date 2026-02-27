package rag

import (
	"database/sql"
	"fmt"
	"strings"
)

func QuerytoAnswer(query string, db *sql.DB) (string, error) {
	contexts, err := SearchSimilar(db, query)
	if err != nil {
		return "", fmt.Errorf("retriever failed to get context %w", err)
	}

	context := strings.Join(contexts, "\n\n")

	result, err := AskLLm(query, context)
	if err != nil {
		return "", fmt.Errorf("llm failed to get result %w", err)
	}
	return result, nil
}
