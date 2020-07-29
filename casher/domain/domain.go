package domain

import "time"

type Item struct {
	ID        int
	Name      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PurchaseHistory struct {
	ID        int
	StoreID   int
	StaffID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Staff struct {
	ID        int
	Name      string
	StoreID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Store struct {
	ID         int
	Branch     string
	Address    string
	CreatedAt  time.Time
	UpdateddAt time.Time
}

type Transaction struct {
	ID         int
	PurchaseID int
	ItemID     int
	Amount     int
	CreatedAt  time.Time
	UpdateddAt time.Time
}
