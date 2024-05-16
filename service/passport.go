package service

import (
	"context"
	"test-task-sw/entity"
)

type PassportService struct {
	passportRepo passportRepo
}

func NewPassportService(repo passportRepo) *PassportService {
	return &PassportService{
		repo,
	}
}

func (p *PassportService) CreatePassport(ctx context.Context, passport entity.Passport) (int64, error) {
	passId, err := p.passportRepo.Create(ctx, passport)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	return passId, nil
}
