package graph

import "github.com/ebalkanski/graphql/graph/model"

type Resolver struct {
	todos      []*model.Todo
	users      []*model.User
	usersChan  chan *model.User
	lastTodoId int
	lastUserId int
}

func NewResolver() *Resolver {
	users := make([]*model.User, 0)
	users = append(users, &model.User{ID: "1", Name: "fphilip"})
	users = append(users, &model.User{ID: "2", Name: "lturanga"})

	return &Resolver{
		todos:      make([]*model.Todo, 0),
		users:      users,
		usersChan:  make(chan *model.User),
		lastUserId: 3,
	}
}
