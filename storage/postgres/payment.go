package postgres

import (
	"context"
	"database/sql"
	pb "payment_service/genproto/payment"
)

type PaymentRepo struct {
	DB *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{DB: db}
}

func (p *PaymentRepo) CreatePayment(ctx context.Context, pay *pb.PaymentDetails) (*pb.Status, error) {
	query := `
	insert into payments (
		reservation_id, amount, method
	)
	values (
		$1, $2, $3
	)`
	_, err := p.DB.ExecContext(ctx, query, pay.ReservationId, pay.Amount, pay.PaymentMethod)
	if err != nil {
		return nil, err
	}

	return &pb.Status{Status: "pending"}, nil
}

func (p *PaymentRepo) GetPayment(ctx context.Context, id string) (*pb.PaymentInfo, error) {
	pay := pb.PaymentInfo{Id: id}
	query := `
	select
		reservation_id, amount, method, status
	from
		payments
	where
		deleted_at IS NULL and id = $1`

	row := p.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&pay.ReservationId, &pay.Amount, &pay.PaymentMethod, &pay.PaymentStatus)
	if err != nil {
		return nil, err
	}

	return &pay, nil
}

func (p *PaymentRepo) UpdatePayment(ctx context.Context, pay *pb.PaymentInfo) error {
	query := `
	update
		payments
	set
		reservation_id = $1,
		amount = $2,
		method = $3,
		status = $4,
		updated_at = NOW()
	where
		deleted_at IS NULL and id = $5`

	res, err := p.DB.ExecContext(ctx, query,
		pay.ReservationId, pay.Amount, pay.PaymentMethod, pay.PaymentStatus, pay.Id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (p *PaymentRepo) SearchByReservationID(ctx context.Context, id string) (*pb.PaymentInfo, error) {
	query := `
	select
		id, amount, method, status
	from
		payments
	where
		deleted_at is null and reservation_id = $1`

	pay := pb.PaymentInfo{ReservationId: id}
	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&pay.Id, &pay.Amount, &pay.PaymentMethod, &pay.PaymentStatus)
	if err != nil {
		return nil, err
	}

	return &pay, nil
}
