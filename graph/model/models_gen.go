// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type Todo struct {
	ID         string `json:"id"`
	Text       string `json:"text"`
	TextLength int    `json:"textLength"`
	Done       bool   `json:"done"`
	User       *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
