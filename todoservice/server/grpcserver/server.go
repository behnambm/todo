package grpcserver

import (
	"context"
	"fmt"
	pb "github.com/behnambm/todo/todoservice/protobuf"
	"github.com/behnambm/todo/todoservice/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type TodoService interface {
	GetUserTodos(int64) ([]types.Todo, error)
	GetTodo(int64) (types.Todo, error)
	GetUserTodosWithItems(int64) ([]types.TodoWithItems, error)
	GetTodoWithItems(int64) (types.TodoWithItems, error)
	GetItem(int64) (types.Item, error)
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
	userId := in.GetUserId()
	userTodos, getTodoErr := s.todoSvc.GetUserTodos(userId)
	if getTodoErr != nil {
		log.Println("[gRPC] GetUserTodos - ", getTodoErr)

		return &pb.UserTodosReply{}, status.Error(codes.NotFound, "user not found")
	}

	var reply []*pb.TodoReply

	for _, todo := range userTodos {
		todoMessage := &pb.TodoReply{Id: int64(todo.ID), Name: todo.Name, Description: todo.Description}
		reply = append(reply, todoMessage)
	}

	return &pb.UserTodosReply{Todos: reply}, nil
}

func (s Server) GetUserTodosWithItems(ctx context.Context, in *pb.UserTodosRequest) (*pb.UserTodosWithItemsReply, error) {
	userId := in.GetUserId()
	userTodosWithItems, getTodoErr := s.todoSvc.GetUserTodosWithItems(userId)
	if getTodoErr != nil {
		log.Println("[gRPC] GetUserTodosWithItems - ", getTodoErr)

		return &pb.UserTodosWithItemsReply{}, status.Error(codes.NotFound, "user not found")
	}

	var reply []*pb.TodoWithItemsReply

	for _, todo := range userTodosWithItems {
		todoMessage := &pb.TodoWithItemsReply{
			Id: todo.ID, Name: todo.Name, Description: todo.Description, UserId: todo.UserId,
		}

		var todoItems []*pb.ItemReply
		for _, item := range todo.Items {
			itemMessage := &pb.ItemReply{Id: item.ID, Title: item.Title, Priority: int64(item.Priority)}
			todoItems = append(todoItems, itemMessage)
		}
		todoMessage.Items = todoItems
		reply = append(reply, todoMessage)
	}

	return &pb.UserTodosWithItemsReply{Todos: reply}, nil
}

func (s Server) GetTodoWithItems(ctx context.Context, in *pb.TodoRequest) (*pb.TodoWithItemsReply, error) {
	todoId := in.GetTodoId()
	todoWithItems, getTodoErr := s.todoSvc.GetTodoWithItems(todoId)
	if getTodoErr != nil {
		log.Println("[gRPC] GetTodoWithItems - ", getTodoErr)

		return &pb.TodoWithItemsReply{}, status.Error(codes.NotFound, "todo not found")
	}

	todoWithItemsMessage := &pb.TodoWithItemsReply{
		Id:          todoWithItems.ID,
		Name:        todoWithItems.Name,
		Description: todoWithItems.Description,
	}

	var itemsMessage []*pb.ItemReply
	for _, item := range todoWithItems.Items {
		itemMessage := &pb.ItemReply{Id: int64(item.ID), Title: item.Title, Priority: int64(item.Priority)}
		itemsMessage = append(itemsMessage, itemMessage)
	}

	todoWithItemsMessage.Items = itemsMessage

	return todoWithItemsMessage, nil
}

func (s Server) GetTodo(ctx context.Context, in *pb.TodoRequest) (*pb.TodoReply, error) {
	todoId := in.GetTodoId()
	todo, err := s.todoSvc.GetTodo(todoId)
	if err != nil {
		log.Println("[gRPC] GetTodo - ", err)

		return &pb.TodoReply{}, status.Error(codes.NotFound, "todo not found")
	}

	return &pb.TodoReply{
		Id:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		UserId:      todo.UserId,
	}, nil
}

func (s Server) GetItem(ctx context.Context, in *pb.ItemRequest) (*pb.ItemReply, error) {
	itemId := in.GetItemId()
	item, err := s.todoSvc.GetItem(itemId)
	if err != nil {
		log.Println("[gRPC] GetItem - ", err)

		return &pb.ItemReply{}, status.Error(codes.NotFound, "item not found")
	}

	return &pb.ItemReply{
		Id:       item.ID,
		Title:    item.Title,
		Priority: int64(item.Priority),
		UserId:   item.UserId,
		TodoId:   item.TodoId,
	}, nil
}
