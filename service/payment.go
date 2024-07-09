package service

import (
	"database/sql"
	pb "payment_service/genproto/payment"
	"payment_service/storage/postgres"
)

type PaymentService struct {
	pb.UnimplementedPaymentServer
	Repo *postgres.PaymentRepo
}

func NewPaymentService(db *sql.DB) *PaymentService {
	return &PaymentService{Repo: postgres.NewPaymentRepo(db)}
}
