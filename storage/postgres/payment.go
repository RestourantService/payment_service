package postgres

import "database/sql"

type PaymentRepo struct {
	DB *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{DB: db}
}
