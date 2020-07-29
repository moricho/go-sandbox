package main

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/moricho/casher/domain"
)

/////// Item Repository ///////
type ItemRepository interface {
	Get(id int) (*domain.Item, error)
}

type ItemRepositoryImpl struct {
	db *sql.DB
}

func (r *ItemRepositoryImpl) Get(id int) (*domain.Item, error) {
	item := &domain.Item{}
	sql, args, err := sq.Select(
		"id",
		"name",
		"price",
	).From("items").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return item, err
	}

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return item, err
	}
	err = stmt.QueryRow(args...).Scan(
		&item.ID,
		&item.Name,
		&item.Price,
	)
	if err != nil {
		return item, err
	}

	return item, nil
}
