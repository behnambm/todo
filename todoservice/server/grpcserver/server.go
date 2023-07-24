package grpcserver

import (
	"context"
	"fmt"
	"github.com/behnambm/todo/todoservice/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"

	pb "github.com/behnambm/todo/todoservice/protobuf"
)

type TodoService interface {
	GetUserTodos(int) ([]types.Todo, error)
	GetUserTodosWithItems(int) ([]types.TodoWithItems, error)
	GetTodoWithItems(int) (types.TodoWithItems, error)
}

// Server is used to implement User RPC service
type Server struct {
	ListenAddr string
	todoSvc    TodoService
	pb.UnimplementedTodoServer
}

func New(addr string, todoService TodoService) *Server {
	return &Server{
		ListenAddr: addr,
		todoSvc:    todoService,
	}
}

func (s Server) Run(ctx context.Context) {
	lis, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	pb.RegisterTodoServer(rpcServer, s)

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

func (s Server) GetUserTodos(ctx context.Context, in *pb.UserTodosRequest) (*pb.UserTodosReply, error) {
	userId, err := strconv.Atoi(in.GetUserId())
	if err != nil {
		log.Printf("[gRPC] GetUserTodos - userId is not valid %s \n", in.GetUserId())

		return &pb.UserTodosReply{}, status.Error(codes.InvalidArgument, "userId is not valid")
	}

	userTodos, getTodoErr := s.todoSvc.GetUserTodos(userId)
	if getTodoErr != nil {
		log.Println("[gRPC] GetUserTodos - ", getTodoErr)

		return &pb.UserTodosReply{}, status.Error(codes.NotFound, "user not found")
	}

	var reply []*pb.TodoMessage

	for _, todo := range userTodos {
		todoMessage := &pb.TodoMessage{Id: int64(todo.ID), Name: todo.Name, Description: todo.Description}
		reply = append(reply, todoMessage)
	}

	return &pb.UserTodosReply{Todos: reply}, nil
}

func (s Server) GetUserTodosWithItems(ctx context.Context, in *pb.UserTodosRequest) (*pb.UserTodosWithItemsReply, error) {
	userId, err := strconv.Atoi(in.GetUserId())
	if err != nil {
		log.Printf("[gRPC] GetUserTodosWithItems - userId is not valid %s \n", in.GetUserId())

		return &pb.UserTodosWithItemsReply{}, status.Error(codes.InvalidArgument, "userId is not valid")
	}

	userTodosWithItems, getTodoErr := s.todoSvc.GetUserTodosWithItems(userId)
	if getTodoErr != nil {
		log.Println("[gRPC] GetUserTodosWithItems - ", getTodoErr)

		return &pb.UserTodosWithItemsReply{}, status.Error(codes.NotFound, "user not found")
	}

	var reply []*pb.TodoWithItemsMessage

	for _, todo := range userTodosWithItems {
		todoMessage := &pb.TodoWithItemsMessage{Id: int64(todo.ID), Name: todo.Name, Description: todo.Description}

		var todoItems []*pb.ItemMessage
		for _, item := range todo.Items {
			itemMessage := &pb.ItemMessage{Id: int64(item.ID), Title: item.Title, Priority: int64(item.Priority)}
			todoItems = append(todoItems, itemMessage)
		}
		todoMessage.Items = todoItems
		reply = append(reply, todoMessage)
	}

	return &pb.UserTodosWithItemsReply{Todos: reply}, nil
}

func (s Server) GetTodoWithItems(ctx context.Context, in *pb.TodoRequest) (*pb.TodoWithItemsMessage, error) {
	todoId, err := strconv.Atoi(in.GetTodoId())
	if err != nil {
		log.Printf("[gRPC] GetTodoWithItems - todoId is not valid %s \n", in.GetTodoId())

		return &pb.TodoWithItemsMessage{}, status.Error(codes.InvalidArgument, "todoId is not valid")
	}

	todoWithItems, getTodoErr := s.todoSvc.GetTodoWithItems(todoId)
	if getTodoErr != nil {
		log.Println("[gRPC] GetTodoWithItems - ", getTodoErr)

		return &pb.TodoWithItemsMessage{}, status.Error(codes.NotFound, "todo not found")
	}

	todoWithItemsMessage := &pb.TodoWithItemsMessage{
		Id:          int64(todoWithItems.ID),
		Name:        todoWithItems.Name,
		Description: todoWithItems.Description,
	}

	var itemsMessage []*pb.ItemMessage
	for _, item := range todoWithItems.Items {
		itemMessage := &pb.ItemMessage{Id: int64(item.ID), Title: item.Title, Priority: int64(item.Priority)}
		itemsMessage = append(itemsMessage, itemMessage)
	}

	todoWithItemsMessage.Items = itemsMessage

	return todoWithItemsMessage, nil
}
