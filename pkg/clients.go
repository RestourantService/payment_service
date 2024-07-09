package pkg

import (
	"errors"
	"log"
	"payment_service/config"
	pb "payment_service/genproto/reservation"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUserClient(cfg config.Config) pb.ReservationClient {
	conn, err := grpc.NewClient(cfg.Server.RESERVATION_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pb.NewReservationClient(conn)
}
