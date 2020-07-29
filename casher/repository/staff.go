package main

import (
	"database/sql"

	"github.com/moricho/casher/domain"
)

/////// Staff Repository ///////
type StaffRepository interface {
	Get(id int) (*domain.Staff, error)
}

type StaffRepositoryImpl struct {
	db *sql.DB
}
