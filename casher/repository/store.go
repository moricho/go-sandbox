package main

import (
	"database/sql"

	"github.com/moricho/casher/domain"
)

/////// Store Repository ///////
type StoreRepository interface {
	Get(id int) (*domain.Store, error)
}

type StoreRepositoryImpl struct {
	db *sql.DB
}
