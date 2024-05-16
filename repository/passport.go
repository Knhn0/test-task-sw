package repository

import (
	"context"
	"test-task-sw/entity"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/repository/query"
)

type PassportRepository struct {
	db *tpostgres.Postgres
}

func NewPassportRepository(db *tpostgres.Postgres) *PassportRepository {
	return &PassportRepository{
		db: db,
	}
}

func (p *PassportRepository) Create(ctx context.Context, passport entity.Passport) (int64, error) {
	passportData := []interface{}{
		passport.Type,
		passport.Number,
	}

	_, err := p.db.ExecContext(ctx, query.InsertPassportData, passportData...)
	if err != nil {
		return 0, err
	}

	var passId int64
	err = p.db.GetContext(ctx, &passId, query.InsertPassportData, passportData...)

	return passId, nil
}
