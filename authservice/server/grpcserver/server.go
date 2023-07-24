package grpcserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "github.com/behnambm/todo/authservice/protobuf"
)

type TokenService interface {
	IsTokenValid(string) bool
	GetToken(int) (string, error)
}

type Server struct {
	ListenAddr string
	tokenSvc   TokenService
	pb.UnimplementedTokenServer
}

func New(addr string, tokenService TokenService) *Server {
	return &Server{
		ListenAddr: addr,
		tokenSvc:   tokenService,
	}
}

func (s Server) Run(ctx context.Context) {
	lis, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		log.Fatalf("[gRPC] failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	pb.RegisterTokenServer(rpcServer, s)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("shutting down...")
			rpcServer.GracefulStop()
			fmt.Println("shutdown complete.")
		}
	}()

	log.Printf("[gRPC] server started on [%s]...\n", s.ListenAddr)
	if rpcErr := rpcServer.Serve(lis); rpcErr != nil {
		log.Fatalf("failed to serve: %v", rpcErr)
	}
}

func (s Server) IsTokenValid(ctx context.Context, in *pb.TokenRequest) (*pb.TokenReply, error) {
	token := in.GetToken()

	if s.tokenSvc.IsTokenValid(token) {
		return &pb.TokenReply{IsValid: true}, nil
	}

	return &pb.TokenReply{IsValid: false}, nil
}

func (s Server) GetToken(ctx context.Context, in *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	userid := in.GetUserid()

	token, err := s.tokenSvc.GetToken(int(userid))
	if err != nil {
		log.Println("[gRPC] GetToken - unable to get token for user ", userid, err)

		return &pb.GetTokenReply{}, status.Error(codes.Unknown, "unable to get token")
	}

	return &pb.GetTokenReply{Token: token}, nil
}
