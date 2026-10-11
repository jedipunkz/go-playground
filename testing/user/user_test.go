package user

import (
	"context"
	"errors"
	"testing"
)

// fakeRepo はテスト用の UserRepository 実装
type fakeRepo struct {
	users map[int]*User
	err   error // 返したいエラーを差し込む
}

func (f *fakeRepo) FindByID(ctx context.Context, id int) (*User, error) {
	if f.err != nil {
		return nil, f.err
	}
	u, ok := f.users[id]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}

func TestGreeting(t *testing.T) {
	errDB := errors.New("connection refused")

	tests := []struct {
		name    string
		repo    *fakeRepo
		id      int
		want    string
		wantErr error
	}{
		{
			name: "存在するユーザー",
			repo: &fakeRepo{users: map[int]*User{1: {ID: 1, Name: "jedi"}}},
			id:   1,
			want: "こんにちは、jediさん",
		},
		{
			name: "存在しないユーザー",
			repo: &fakeRepo{users: map[int]*User{}},
			id:   2,
			want: "こんにちは、ゲストさん",
		},
		{
			name:    "DB エラー",
			repo:    &fakeRepo{err: errDB},
			id:      1,
			wantErr: errDB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.repo)
			got, err := s.Greeting(t.Context(), tt.id)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
