package graph

import (
	"github.com/ebalkanski/graphql/graph/model"
	pb "github.com/ebalkanski/grpc/proto"
)

type Resolver struct {
	grpc       pb.UsersClient
	todos      []*model.Todo
	lastTodoId int
	usersChan  chan *model.User
}

func NewResolver(grpc pb.UsersClient) *Resolver {
	return &Resolver{
		grpc:      grpc,
		todos:     make([]*model.Todo, 0),
		usersChan: make(chan *model.User),
	}
}
