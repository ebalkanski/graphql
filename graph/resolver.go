package graph

import (
	"context"

	"github.com/ebalkanski/graphql/graph/model"
)

type Client interface {
	Users(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, u *model.User) (*model.User, error)
	Close() error
}

type Resolver struct {
	client     Client
	todos      []*model.Todo
	lastTodoId int
	usersChan  chan *model.User
}

func NewResolver(client Client) *Resolver {
	return &Resolver{
		client:    client,
		todos:     make([]*model.Todo, 0),
		usersChan: make(chan *model.User),
	}
}
