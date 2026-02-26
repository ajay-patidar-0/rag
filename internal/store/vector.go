package store

import "database/sql"

type VectorDB struct {
	db *sql.DB
}

func NewVectorDB(db *sql.DB) *VectorDB {
	return &VectorDB{
		db: db,
	}
}
