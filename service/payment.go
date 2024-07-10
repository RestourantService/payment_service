package service

import (
	"context"
	"database/sql"
	pb "payment_service/genproto/payment"
	pbr "payment_service/genproto/reservation"
	"payment_service/storage/postgres"

	"github.com/pkg/errors"
)

type PaymentService struct {
	pb.UnimplementedPaymentServer
	Repo              *postgres.PaymentRepo
	ReservationClient pbr.ReservationClient
}

func NewPaymentService(db *sql.DB, reservation pbr.ReservationClient) *PaymentService {
	return &PaymentService{
		Repo:              postgres.NewPaymentRepo(db),
		ReservationClient: reservation,
	}
}

func (p *PaymentService) CreatePayment(ctx context.Context, req *pb.PaymentDetails) (*pb.Status, error) {
	status, err := p.ReservationClient.ValidateReservation(ctx, &pbr.ID{Id: req.ReservationId})
	if err != nil {
        return nil, errors.Wrap(err, "failed to validate reservation")
    }
    if !status.Successful {
        return nil, errors.New("reservation validation failed")
    }

	resp, err := p.Repo.CreatePayment(ctx, req)
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

func (p *PaymentService) SearchByReservationID(ctx context.Context, req *pb.ID) (*pb.PaymentInfo, error) {
	resp, err := p.Repo.SearchByReservationID(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find payment")
	}

	return resp, nil
}
