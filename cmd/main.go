package main

import (
	"log"
	"net"
	"payment_service/config"
	pb "payment_service/genproto/payment"
	"payment_service/pkg"
	"payment_service/service"
	"payment_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	lis, err := net.Listen("tcp", cfg.Server.PAYMENT_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	reservationClient := pkg.CreateReservationClient(*cfg)
	paymentService := service.NewPaymentService(db, reservationClient)
	server := grpc.NewServer()
	pb.RegisterPaymentServer(server, paymentService)

	log.Printf("server listening at %v", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
