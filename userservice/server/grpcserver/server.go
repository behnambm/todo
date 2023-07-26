package grpcserver

import (
	"context"
	"fmt"
	"github.com/behnambm/todo/userservice/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "github.com/behnambm/todo/userservice/protobuf"
)

type UserService interface {
	GetUserByEmail(string) (types.User, error)
	GetUserById(int64) (types.User, error)
}

// Server is used to implement User RPC service
type Server struct {
	ListenAddr string
	userSvc    UserService
	pb.UnimplementedUserServer
}

func New(addr string, userServer UserService) *Server {
	return &Server{
		ListenAddr: addr,
		userSvc:    userServer,
	}
}

func (s Server) Run(ctx context.Context) {
	lis, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	pb.RegisterUserServer(rpcServer, s)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("shutting down...")
			rpcServer.GracefulStop()
			fmt.Println("shutdown complete.")
		}
	}()

	log.Printf("[gRPC] server started on [%s]...\n", s.ListenAddr)

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s Server) GetUserByEmail(ctx context.Context, in *pb.GetUserByEmailRequest) (*pb.GetUserReply, error) {
	user, err := s.userSvc.GetUserByEmail(in.GetEmail())
	if err != nil {
		log.Printf("[gRPC] GetUserByEmail - user [%s] not found - %v", in.GetEmail(), err)

		return &pb.GetUserReply{}, status.Error(codes.NotFound, "user not found")
	}

	return &pb.GetUserReply{Name: user.Name, Email: user.Email, Id: user.ID, Password: user.Password}, nil
}

func (s Server) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserReply, error) {
	user, err := s.userSvc.GetUserById(in.GetUserId())
	if err != nil {
		log.Printf("[gRPC] GetUserById - user [%s] not found - %v", in.GetUserId(), err)

		return &pb.GetUserReply{}, status.Error(codes.NotFound, "user not found")
	}

	return &pb.GetUserReply{Name: user.Name, Email: user.Email, Id: user.ID, Password: user.Password}, nil
}
