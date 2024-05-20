package department

import (
	"context"
	"database/sql"
	depqueries "test-task-sw/database/department/queries"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/service/models"
)

type Repo interface {
	Create(ctx context.Context, tx *sql.Tx, department models.Department) (int64, error)
	Delete(ctx context.Context, departmentId int64) error
	Update(ctx context.Context, tx *sql.Tx, departmentId int64, department models.Department) error
}

type Impl struct {
	db *tpostgres.Postgres
}

func NewRepository(db *tpostgres.Postgres) Repo {
	return &Impl{
		db: db,
	}
}

func (d *Impl) Create(ctx context.Context, tx *sql.Tx, department models.Department) (int64, error) {
	var depId int64

	err := tx.QueryRowContext(ctx, depqueries.CreateDepartment, department.Name, department.Phone).Scan(&depId)
	if err != nil {
		return 0, err
	}

	return depId, nil
}

func (d *Impl) Delete(ctx context.Context, departmentId int64) error {
	_, err := d.db.ExecContext(ctx, depqueries.DeleteDepartment, departmentId)
	if err != nil {
		return err
	}

	return nil
}

func (d *Impl) Update(ctx context.Context, tx *sql.Tx, departmentId int64, department models.Department) error {
	_, err := tx.ExecContext(ctx, depqueries.UpdateDepartment, departmentId, department.Phone, department.Name)
	return err
}
