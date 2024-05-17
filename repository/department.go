package repository

import (
	"context"
	"test-task-sw/entity"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/repository/query"
)

type DepartmentRepository struct {
	db *tpostgres.Postgres
}

func NewDepartmentRepository(db *tpostgres.Postgres) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (d *DepartmentRepository) Create(ctx context.Context, department entity.Department) (int64, error) {
	departmentData := []interface{}{
		department.Name,
		department.Phone,
	}

	var depId int64
	err := d.db.GetContext(ctx, &depId, query.InsertDepartmentData, departmentData...)
	if err != nil {
		return 0, err
	}

	return depId, nil
}

func (d *DepartmentRepository) Delete(ctx context.Context, departmentId int64) error {
	_, err := d.db.ExecContext(ctx, query.DeleteDepartment, departmentId)
	if err != nil {
		return err
	}
	return nil
}
