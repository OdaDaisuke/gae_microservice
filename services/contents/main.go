package main

import (
	"context"
	pb "github.com/OdaDaisuke/gae_sand/pb/content"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ContentServer struct {
}

func (s *ContentServer) GetContent(c context.Context, r *pb.GetContentRequest) (*pb.GetContentResponse, error) {
	cr := &pb.GetContentResponse{
		Content: &pb.Content{
			Id: r.Id,
			Name: "aa",
			Price: "1",
			Description: "",
		},
	}
	return cr, nil
}

func (s *ContentServer) GetContents(c context.Context, r *pb.GetContentsRequest) (*pb.GetContentsResponse, error) {
	contents := []*pb.Content{}
	return &pb.GetContentsResponse{
		Contents: contents,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal("failed to listen")
	}
	s := grpc.NewServer()
	pb.RegisterContentServiceServer(s, &ContentServer{})
	log.Println("Server listening on port: 19003")
	if err := s.Serve(lis); err != nil {
		log.Fatal("err")
	}
}