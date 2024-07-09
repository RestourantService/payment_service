package service

import (
	"context"
	"database/sql"
	pb "payment_service/genproto/payment"
	"payment_service/storage/postgres"

	"github.com/pkg/errors"
)

type PaymentService struct {
	pb.UnimplementedPaymentServer
	Repo *postgres.PaymentRepo
}

func NewPaymentService(db *sql.DB) *PaymentService {
	return &PaymentService{Repo: postgres.NewPaymentRepo(db)}
}

func (p *PaymentService) MakePayment(ctx context.Context, req *pb.PaymentDetails) (*pb.Status, error) {
	resp, err := p.Repo.MakePayment(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make payment")
	}

	return resp, nil
}

func (p *PaymentService) GetPayment(ctx context.Context, req *pb.ID) (*pb.PaymentInfo, error) {
	resp, err := p.Repo.GetPayment(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read payment")
	}

	return resp, nil
}

func (p *PaymentService) UpdatePayment(ctx context.Context, req *pb.PaymentInfo) (*pb.Void, error) {
	err := p.Repo.UpdatePayment(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update payment")
	}

	return &pb.Void{}, nil
}
