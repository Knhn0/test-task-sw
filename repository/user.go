package repository

import (
	"context"
	"test-task-sw/entity"
	"test-task-sw/lib/tpostgres"
)

type UserRepository struct {
	db *tpostgres.Postgres
}

func NewUserRepository(db *tpostgres.Postgres) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUser(ctx context.Context, userId int) (entity.User, error) {
	return entity.User{}, nil
}
func (u *UserRepository) Create(ctx context.Context, user entity.User, passportHash string) (int, error) {
	return 1, nil
}
