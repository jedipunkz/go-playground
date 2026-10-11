package user

import (
	"context"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("user not found")

type User struct {
	ID   int
	Name string
}

// UserRepository はユーザーの永続化を抽象化する
type UserRepository interface {
	FindByID(ctx context.Context, id int) (*User, error)
}

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

// Greeting はユーザー名入りの挨拶文を返す
func (s *Service) Greeting(ctx context.Context, id int) (string, error) {
	u, err := s.repo.FindByID(ctx, id)
	if errors.Is(err, ErrNotFound) {
		return "こんにちは、ゲストさん", nil
	}
	if err != nil {
		return "", fmt.Errorf("find user %d: %w", id, err)
	}
	return fmt.Sprintf("こんにちは、%sさん", u.Name), nil
}
