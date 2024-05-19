package service

import (
	"context"
	"test-task-sw/service/models"
)

type PassportService struct {
	passportRepo passportRepo
}

func NewPassportService(repo passportRepo) *PassportService {
	return &PassportService{
		repo,
	}
}

func (p *PassportService) CreatePassport(ctx context.Context, passport models.Passport) (int64, error) {
	passId, err := p.passportRepo.Create(ctx, passport)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	return passId, nil
}

func (p *PassportService) DeletePassport(ctx context.Context, passportId int64) error {
	err := p.passportRepo.Delete(ctx, passportId)
	switch {
	case err == nil:
	default:
		return err
	}

	return nil
}
