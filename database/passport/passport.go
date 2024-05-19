package passport

import (
	"context"
	passqueries "test-task-sw/database/passport/queries"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/service/models"
)

type PassportRepository struct {
	db *tpostgres.Postgres
}

func NewPassportRepository(db *tpostgres.Postgres) *PassportRepository {
	return &PassportRepository{
		db: db,
	}
}

func (p *PassportRepository) Create(ctx context.Context, passport models.Passport) (int64, error) {
	passportData := []interface{}{
		passport.Type,
		passport.Number,
	}

	var passId int64
	err := p.db.GetContext(ctx, &passId, passqueries.CreatePassport, passportData...)
	if err != nil {
		return 0, err
	}

	return passId, nil
}

func (p *PassportRepository) Delete(ctx context.Context, passportId int64) error {
	_, err := p.db.ExecContext(ctx, passqueries.DeletePassport, passportId)
	if err != nil {
		return err
	}

	return nil
}
