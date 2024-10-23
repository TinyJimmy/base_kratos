package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID    string
	Name  string
	Age   int
	Email string
	Phone string
}

type UserRepo interface {
	// db
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id string, user *User) error
	DeleteUser(ctx context.Context, id string) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Get(ctx context.Context, id string) (p *User, err error) {
	p, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return p, nil
}

func (uc *UserUsecase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) Update(ctx context.Context, id string, user *User) error {
	return uc.repo.UpdateUser(ctx, id, user)
}

func (uc *UserUsecase) Delete(ctx context.Context, id string) error {
	return uc.repo.DeleteUser(ctx, id)
}
