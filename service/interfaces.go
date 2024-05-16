package service

import (
	"context"
	"test-task-sw/entity"
)

type userRepo interface {
	GetUser(ctx context.Context, userId int) (entity.User, error)
	Create(ctx context.Context, user entity.User, passportHash string) (int, error)
}
