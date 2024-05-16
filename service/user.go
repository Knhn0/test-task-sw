package service

import (
	"context"
	"database/sql"
	"errors"
	"test-task-sw/entity"
)

type UserService struct {
	userRepo   userRepo
	dataHasher DataHasher
}

func NewUserService(repo userRepo, dataHasher DataHasher) *UserService {
	return &UserService{
		repo,
		dataHasher,
	}
}

func (u *UserService) Create(ctx context.Context, user entity.User) (int, error) {
	isUserExists, err := u.isUserExists(ctx, user.Id)
	if err != nil {
		return 0, err
	}

	if isUserExists {
		err = ErrAlreadyExists
		return 0, err
	}

	passportHash := ""
	userId, err := u.userRepo.Create(ctx, user, passportHash)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	return userId, nil
}

func (u *UserService) isUserExists(ctx context.Context, userId int) (bool, error) {

	_, err := u.userRepo.GetUser(ctx, userId)

	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return false, nil
	default:
		return false, err
	}

	return true, nil
}
