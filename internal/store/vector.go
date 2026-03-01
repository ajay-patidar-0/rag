package store

import (
	"database/sql"
	"fmt"
	"strings"
)

type VectorDB struct {
	DB *sql.DB
}

func NewVectorDB(db *sql.DB) *VectorDB {
	return &VectorDB{
		DB: db,
	}
}

func (v *VectorDB) AddVector(content string, vector string) error {
	query := `INSERT INTO documents (content,embedding)
	VALUES($1, $2)`

	_, err := v.DB.Query(query, content, vector)
	if err != nil {
		return fmt.Errorf("Enable to insert embedding %w", err)
	}

	return nil
}

func ToPGVector(vec []float32) string {
	var sb strings.Builder
	sb.WriteString("[")

	for i, v := range vec {
		sb.WriteString(fmt.Sprintf("%f", v))
		if i != len(vec)-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
