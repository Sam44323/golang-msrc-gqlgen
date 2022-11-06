package store

import (
	"context"
	"net/http"

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

// WithStore middleware - injects a store into the context
func WithStore(store *Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store := NewStore()
		ctx := context.WithValue(r.Context(), "STORE", store) // adding store to context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetStoreFromContext - retrieves a store from the context
