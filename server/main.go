package main

import (
	"context"
	"log"
	"net"

	pb "github.com/JoseTheodoro/grpc-go/person"
	"github.com/JoseTheodoro/grpc-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *Server) GetPerson(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	p := &pb.Person{
		Name: in.GetName(),
		Age:  in.GetAge(),
		City: in.GetCity(),
	}
	return &pb.PersonResponse{
		Person: p,
	}, nil
}

func (s *Server) CreatePerson(ctx context.Context, input *pb.PersonRequest) (*pb.PersonResponse, error) {

	p := services.CreatePerson(input)

	pbb := &pb.Person{
		ID:   p.ID,
		Name: p.Name,
		Age:  p.Age,
		City: p.City,
	}

	return &pb.PersonResponse{
		Person: pbb,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &Server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
