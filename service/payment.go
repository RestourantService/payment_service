package service

import (
	"context"
	"database/sql"
	"log/slog"
	pb "payment_service/genproto/payment"
	pbr "payment_service/genproto/reservation"
	"payment_service/pkg/logger"
	"payment_service/storage/postgres"

	"github.com/pkg/errors"
)

type PaymentService struct {
	pb.UnimplementedPaymentServer
	Repo              *postgres.PaymentRepo
	ReservationClient pbr.ReservationClient
	Logger            *slog.Logger
}

func NewPaymentService(db *sql.DB, reservation pbr.ReservationClient) *PaymentService {
	return &PaymentService{
		Repo:              postgres.NewPaymentRepo(db),
		ReservationClient: reservation,
		Logger:            logger.NewLogger(),
	}
}

func (p *PaymentService) CreatePayment(ctx context.Context, req *pb.PaymentDetails) (*pb.Status, error) {
	p.Logger.Info("CreatePayment method is starting")

	status, err := p.ReservationClient.ValidateReservation(ctx, &pbr.ID{Id: req.ReservationId})
	if err != nil {
		err := errors.Wrap(err, "failed to validate reservation")
		p.Logger.Error(err.Error())
		return nil, err
	}
	if !status.Successful {
		err := errors.New("reservation validation failed")
		p.Logger.Error(err.Error())
		return nil, err
	}

	p.Logger.Info("Reservation has been validated")

	resp, err := p.Repo.CreatePayment(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to make payment")
		p.Logger.Error(err.Error())
		return nil, err
	}

	p.Logger.Info("CreatePayment has successfully finished")
	return resp, nil
}

func (p *PaymentService) GetPayment(ctx context.Context, req *pb.ID) (*pb.PaymentInfo, error) {
	p.Logger.Info("GetPayment method is starting")

	resp, err := p.Repo.GetPayment(ctx, req.Id)
	if err != nil {
		err := errors.Wrap(err, "failed to read payment")
		p.Logger.Error(err.Error())
		return nil, err
	}

	p.Logger.Info("GetPayment has successfully finished")
	return resp, nil
}

func (p *PaymentService) UpdatePayment(ctx context.Context, req *pb.PaymentInfo) (*pb.Void, error) {
	p.Logger.Info("UpdatePayment method is starting")

	err := p.Repo.UpdatePayment(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to update payment")
		p.Logger.Error(err.Error())
		return nil, err
	}

	p.Logger.Info("UpdatePayment has successfully finished")
	return &pb.Void{}, nil
}

func (p *PaymentService) SearchByReservationID(ctx context.Context, req *pb.ID) (*pb.PaymentInfo, error) {
	p.Logger.Info("SearchByReservationID method is starting")

	resp, err := p.Repo.SearchByReservationID(ctx, req.Id)
	if err != nil {
		err := errors.Wrap(err, "failed to find payment")
		p.Logger.Error(err.Error())
		return nil, err
	}

	p.Logger.Info("SearchByReservationID has successfully finished")
	return resp, nil
}
