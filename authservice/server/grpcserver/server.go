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
	GetToken(int64) (string, error)
	IsValidWithClaim(string) (map[string]string, bool)
}

type Server struct {
	ListenAddr string
	tokenSvc   TokenService
	pb.UnimplementedAuthServer
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
	pb.RegisterAuthServer(rpcServer, s)

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

func (s Server) GetToken(ctx context.Context, in *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	userId := in.GetUserid()

	token, err := s.tokenSvc.GetToken(userId)
	if err != nil {
		return &pb.GetTokenReply{}, status.Error(codes.Unauthenticated, "unable to get token")
	}

	return &pb.GetTokenReply{Token: token}, nil
}

func (s Server) IsTokenValid(ctx context.Context, in *pb.TokenRequest) (*pb.TokenReply, error) {
	token := in.GetToken()

	if s.tokenSvc.IsTokenValid(token) {
		return &pb.TokenReply{IsValid: true}, nil
	}

	return &pb.TokenReply{IsValid: false}, nil
}

func (s Server) ValidateTokenWithClaims(ctx context.Context, in *pb.ValidateTokenWithClaimsRequest) (*pb.ValidateTokenWithClaimsReply, error) {
	token := in.GetToken()

	claims, ok := s.tokenSvc.IsValidWithClaim(token)
	if !ok {
		return &pb.ValidateTokenWithClaimsReply{Claims: claims, IsValid: false}, nil
	}

	return &pb.ValidateTokenWithClaimsReply{Claims: claims, IsValid: true}, nil
}
