package main

import (
	"database/sql"

	"github.com/moricho/casher/domain"
)

/////// PurchaseHistory Repository ///////
type PurchaseHistoryRepository interface {
	Get(id int) (*domain.PurchaseHistory, error)
}

type PurchaseHistoryRepositoryImpl struct {
	db *sql.DB
}
