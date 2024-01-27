package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
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
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error)
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
		return cmt, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(
	ctx context.Context,
	ID string,
	updatedComment Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedComment)
	if err != nil {
		fmt.Println("error updating comment in service")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return insertedCmt, nil
}
