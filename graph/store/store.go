package store

import (
	"github.com/Sam44323/golang-micro/graph/model"
)

type Store struct {
	Todos []*model.Todo
}

func NewStore() *Store {
	todos := make([]*model.Todo, 0)
	return &Store{
		Todos: todos,
	}
}
