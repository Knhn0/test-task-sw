package passport

import (
	"context"
	"database/sql"
	passqueries "test-task-sw/database/passport/queries"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/service/models"
)

type Repo interface {
	Create(ctx context.Context, tx *sql.Tx, passport models.Passport) (int64, error)
	Delete(ctx context.Context, tx *sql.Tx, passportId int64) error
	Update(ctx context.Context, tx *sql.Tx, passportId int64, passport models.Passport) error
}

type Impl struct {
	db *tpostgres.Postgres
}

func NewRepository(db *tpostgres.Postgres) Repo {
	return &Impl{
		db: db,
	}
}

func (p *Impl) Create(ctx context.Context, tx *sql.Tx, passport models.Passport) (int64, error) {
	var passId int64

	err := tx.QueryRowContext(ctx, passqueries.CreatePassport, passport.Type, passport.Number).Scan(&passId)
	if err != nil {
		return 0, err
	}

	return passId, nil
}

func (p *Impl) Delete(ctx context.Context, tx *sql.Tx, passportId int64) error {
	_, err := tx.ExecContext(ctx, passqueries.DeletePassport, passportId)
	return err
}

func (p *Impl) Update(ctx context.Context, tx *sql.Tx, passportId int64, passport models.Passport) error {
	_, err := tx.ExecContext(ctx, passqueries.UpdatePassport, passportId, passport.Type, passport.Number)
	return err
}
