package comment

import (
	"context"
	"errors"
	"fmt"	
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrUpdateComment = errors.New("failed to update comment by id")
	ErrCreateComment = errors.New("failed to add comment")
	ErrDetetingComment = errors.New("failed to delete comment by id")	
)

type Comment struct {
	ID			string
	Slug 		string
	Body 		string
	Author	string
}

// Repository
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

// Service
type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service {
		Store: store,
	}
}

// Service methods
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	updatedCmt, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrUpdateComment
	}
	
	return updatedCmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return ErrDetetingComment
	}

	return nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrCreateComment
	}
	
	return insertedCmt, nil
}