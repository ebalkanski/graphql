package clients

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/ebalkanski/graphql/graph/model"
	pb "github.com/ebalkanski/grpc/proto"
)

type Client struct {
	users pb.UsersClient
	conn  *grpc.ClientConn
}

func New(addr string) (*Client, error) {
	if addr == "" {
		return nil, fmt.Errorf("missing grpc address")
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("can not connect to grpc: %v", err)
	}

	users := pb.NewUsersClient(conn)

	return &Client{
		users: users,
		conn:  conn,
	}, nil
}

func (c *Client) Users(ctx context.Context) ([]*model.User, error) {
	resp, err := c.users.GetUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		return nil, fmt.Errorf("can not get users from grpc: %v", err)
	}

	var users []*model.User
	for _, u := range resp.GetUsers() {
		users = append(users, &model.User{
			ID:   string(u.GetId()),
			Name: u.GetName(),
		})
	}

	return users, nil
}

func (c *Client) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
	resp, err := c.users.CreateUser(ctx, &pb.CreateUserRequest{
		Name: u.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("can not create user via grpc: %v", err)
	}

	return &model.User{
		ID:   string(resp.GetUser().GetId()),
		Name: resp.GetUser().GetName(),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
