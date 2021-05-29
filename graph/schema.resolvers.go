package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ebalkanski/graphql/graph/generated"
	"github.com/ebalkanski/graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var targetUser *model.User

	for _, user := range r.users {
		if user.ID == input.UserID {
			targetUser = user
			break
		}
	}

	if targetUser == nil {
		return nil, fmt.Errorf("user with id='%s' not found", input.UserID)
	}

	newTodo := &model.Todo{
		ID:   strconv.Itoa(r.lastTodoId),
		Text: input.Text,
		Done: false,
		User: targetUser,
	}
	r.todos = append(r.todos, newTodo)
	r.lastTodoId++

	return newTodo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	for _, user := range r.users {
		if user.Name == input.Name {
			return nil, fmt.Errorf("user with name='%s' already exists", input.Name)
		}
	}

	newUser := &model.User{
		ID:   strconv.Itoa(r.lastUserId),
		Name: input.Name,
	}
	r.users = append(r.users, newUser)
	r.lastUserId++
	r.usersChan <- newUser

	return newUser, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *subscriptionResolver) UserAdded(ctx context.Context) (<-chan *model.User, error) {
	return r.usersChan, nil
}

func (r *todoResolver) TextLength(ctx context.Context, obj *model.Todo) (int, error) {
	return len(obj.Text), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
