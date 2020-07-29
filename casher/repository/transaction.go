package main

import (
	"database/sql"

	"github.com/moricho/casher/domain"
)

/////// Transaction Repository ///////
type TransactionRepository interface {
	Get(id int) (*domain.Transaction, error)
}

type TransactionRepositoryImpl struct {
	db *sql.DB
}
