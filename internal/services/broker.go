package services

import (
	"context"
	ssov1 "github.com/shoumq/sso_protos/gen/go/sso"
	"google.golang.org/grpc"
	"log"
)

type Broker struct {
	conn     ssov1.AuthClient
	grpcConn *grpc.ClientConn
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New() *Broker {
	conn, err := grpc.Dial("localhost:44044", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &Broker{
		conn:     ssov1.NewAuthClient(conn),
		grpcConn: conn,
	}
}

func (b *Broker) Login(request AuthRequest) (string, error) {
	loginReq := &ssov1.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
		AppId:    1,
	}

	loginRes, err := b.conn.Login(context.Background(), loginReq)
	if err != nil {
		return "", err
	}
	return loginRes.Token, nil
}

func (b *Broker) Register(request AuthRequest) (int64, error) {
	regReq := &ssov1.RegisterRequest{
		Email:    request.Email,
		Password: request.Password,
	}

	loginRes, err := b.conn.Register(context.Background(), regReq)
	if err != nil {
		return 0, err
	}
	return loginRes.UserId, nil
}

func (b *Broker) IsAdmin() {
	registerReq := &ssov1.RegisterRequest{
		Email:    "example@example.com",
		Password: "securepassword",
	}

	registerRes, err := b.conn.Register(context.Background(), registerReq)

	isAdminReq := &ssov1.IsAdminRequest{
		UserId: registerRes.UserId,
	}

	isAdminRes, err := b.conn.IsAdmin(context.Background(), isAdminReq)
	if err != nil {
		log.Fatalf("could not check admin status: %v", err)
	}
	log.Printf("Is user admin? %v", isAdminRes.IsAdmin)
}

func (b *Broker) Close() {
	if b.grpcConn != nil {
		b.grpcConn.Close()
	}
}
