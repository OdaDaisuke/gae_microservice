package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	pb "github.com/OdaDaisuke/gae_sand/pb/account"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var startUpTime time.Time
var client *storage.Client

type AccountServer struct {
}

func (s *AccountServer) Signin(c context.Context, r *pb.SigninRequest) (*pb.SigninResponse, error) {
	fmt.Printf("requested user : %s\n", r.User)
	token := fmt.Sprintf("sample token %s %s", r.User, r.Password)
	sr := &pb.SigninResponse{
		Token: token,
	}
	return sr, nil
}

// アカウントサービス
func main() {
	lis, err := net.Listen("tcp", ":19004")
	if err != nil {
		log.Fatal("failed to listen")
	}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &AccountServer{})
	log.Println("Server listening on port: 19004")
	if err := s.Serve(lis); err != nil {
	}
}
