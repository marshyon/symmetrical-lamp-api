package comment

import (
	"context"
	"fmt"
)

// Comment - represents a comment
// structure for this package
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all methods
// that the service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - the struct on which all the
// logic will be built on
type Service struct {
	Store Store
}

// NewService - returns a new comment service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("returning comment from service")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return cmt, err
	}

	return Comment{}, nil
}
