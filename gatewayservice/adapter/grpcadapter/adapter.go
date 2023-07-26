package grpcadapter

import (
	"context"
	"fmt"
	"github.com/behnambm/todo/gatewayservice/protobuf/authproto"
	"github.com/behnambm/todo/gatewayservice/protobuf/todoproto"
	"github.com/behnambm/todo/gatewayservice/protobuf/userproto"
	"github.com/behnambm/todo/gatewayservice/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GRPCAdapter struct {
	authClient authproto.AuthClient
	userClient userproto.UserClient
	todoClient todoproto.TodoClient
}

func New(UserGRPCServiceURL, AuthGRPCServiceURL, TodoGRPCServiceURL string) GRPCAdapter {
	userConn, err := grpc.Dial(UserGRPCServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("user gRPC connection error %v \n", err)
	}
	userClient := userproto.NewUserClient(userConn)

	authConn, err := grpc.Dial(AuthGRPCServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("auth gRPC connection error %v \n", err)
	}
	authClient := authproto.NewAuthClient(authConn)

	todoConn, err := grpc.Dial(TodoGRPCServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("todo gRPC connection error %v \n", err)
	}
	todoClient := todoproto.NewTodoClient(todoConn)

	return GRPCAdapter{
		userClient: userClient,
		authClient: authClient,
		todoClient: todoClient,
	}
}

func (g GRPCAdapter) GetUserByEmail(email string) (types.User, error) {
	user, err := g.userClient.GetUserByEmail(context.Background(), &userproto.GetUserByEmailRequest{Email: email})
	if err != nil {
		return types.User{}, fmt.Errorf("[gRPC ADAPTER] GetUserByEmail - %w", err)
	}

	log.Println("i got the user from gRPC : ", user)

	return types.User{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		ID:       user.GetId(),
	}, nil
}

func (g GRPCAdapter) GetToken(userId int64) (string, error) {
	token, err := g.authClient.GetToken(
		context.Background(), &authproto.GetTokenRequest{Userid: userId},
	)
	if err != nil {
		return "", fmt.Errorf("[gRPC ADAPTER] GetToken - %w", err)
	}

	return token.Token, nil
}

func (g GRPCAdapter) IsValidWithClaim(token string) (map[string]string, bool) {
	reply, err := g.authClient.ValidateTokenWithClaims(context.Background(), &authproto.ValidateTokenWithClaimsRequest{
		Token: token,
	})
	if err != nil {
		return nil, false
	}

	return reply.Claims, reply.IsValid
}

func (g GRPCAdapter) GetUserByID(userId int64) (types.User, error) {
	user, err := g.userClient.GetUserById(context.Background(), &userproto.GetUserByIdRequest{UserId: userId})
	if err != nil {
		return types.User{}, fmt.Errorf("[gRPC] GetUserByID - %w", err)
	}

	return types.User{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		ID:       user.GetId(),
	}, nil
}

func (g GRPCAdapter) GetUserTodosWithItems(userId int64) ([]types.TodoWithItems, error) {
	todosReply, err := g.todoClient.GetUserTodosWithItems(
		context.Background(), &todoproto.UserTodosRequest{UserId: userId},
	)
	if err != nil {
		return nil, fmt.Errorf("[gRPC] GetUserTodosWithItems - %w", err)
	}

	var todosWithItems []types.TodoWithItems

	for _, todoReply := range todosReply.Todos {
		var items []types.MinimalItem

		for _, itemReply := range todoReply.Items {
			items = append(items, types.MinimalItem{
				ID:       itemReply.GetId(),
				Title:    itemReply.GetTitle(),
				Priority: int(itemReply.GetPriority()),
			})
		}

		todosWithItems = append(todosWithItems, types.TodoWithItems{
			Todo: types.Todo{
				ID:          todoReply.GetId(),
				Name:        todoReply.GetName(),
				Description: todoReply.GetDescription(),
				UserId:      todoReply.GetUserId(),
			},
			Items: items,
		})
	}

	return todosWithItems, nil
}

func (g GRPCAdapter) GetTodo(todoId int64) (types.Todo, error) {
	todoReply, err := g.todoClient.GetTodo(context.Background(), &todoproto.TodoRequest{TodoId: todoId})
	if err != nil {
		return types.Todo{}, fmt.Errorf("[gRPC] GetTodo - %w", err)
	}

	return types.Todo{
		ID:          todoReply.GetId(),
		Name:        todoReply.GetName(),
		Description: todoReply.GetDescription(),
		UserId:      todoReply.GetUserId(),
	}, nil
}

func (g GRPCAdapter) GetItem(itemId int64) (types.Item, error) {
	reply, err := g.todoClient.GetItem(context.Background(), &todoproto.ItemRequest{ItemId: itemId})
	if err != nil {
		return types.Item{}, fmt.Errorf("[gRPC] GetItem - %w", err)
	}

	return types.Item{
		ID:       reply.GetId(),
		Title:    reply.GetTitle(),
		Priority: int(reply.GetPriority()),
		UserId:   reply.GetUserId(),
		TodoId:   reply.GetTodoId(),
	}, nil
}
