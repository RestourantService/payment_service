package postgres

import (
	"context"
	pb "payment_service/genproto/payment"
	"reflect"
	"testing"
)

func TestMakePayment(t *testing.T) {
	db, _ := ConnectDB()
	defer db.Close()

	repo := NewPaymentRepo(db)

	pay := &pb.PaymentDetails{
		ReservationId: "550e8400-e29b-41d4-a716-446655444001",
		Amount:        100.00,
		PaymentMethod: "cash",
	}

	_, err := repo.CreatePayment(context.Background(), pay)
	if err != nil {
		t.Error("Error making payment")
	}
}

func TestGetPayment(t *testing.T) {
	db, _ := ConnectDB()
	defer db.Close()

	repo := NewPaymentRepo(db)

	id := "a1234567-89ab-cdef-0123-456789abcdef"

	pay, err := repo.GetPayment(context.Background(), id)
	if err != nil {
		t.Error("Error getting payment")
	}

	exp := pb.PaymentInfo{
		Id:            "a1234567-89ab-cdef-0123-456789abcdef",
		ReservationId: "550e8400-e29b-41d4-a716-446655440001",
		Amount:        100.00,
		PaymentMethod: "cash",
		PaymentStatus: "completed",
	}
	if !reflect.DeepEqual(&exp, pay) {
		t.Error("Payment ID mismatch")
	}
}

func TestUpdatePayment(t *testing.T) {
	db, _ := ConnectDB()
	defer db.Close()

	repo := NewPaymentRepo(db)

	id := "550e8400-e29b-41d4-a716-446655440001"

	pay := &pb.PaymentInfo{
		Id:            id,
		ReservationId: "550e8400-e29b-41d4-a716-446655440001",
		Amount:        100.00,
		PaymentMethod: "cash",
		PaymentStatus: "completed",
	}

	err := repo.UpdatePayment(context.Background(), pay)
	if err != nil {
		t.Error("Error updating payment")
	}
}
