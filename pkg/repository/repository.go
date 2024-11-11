package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}
type FinanceList interface {
}

type FinanceItem interface {
}

type Repository struct {
	Authorization
	FinanceList
	FinanceItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
