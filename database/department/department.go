package department

import (
	"context"
	depqueries "test-task-sw/database/department/queries"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/service/models"
)

type DepartmentRepository struct {
	db *tpostgres.Postgres
}

func NewDepartmentRepository(db *tpostgres.Postgres) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (d *DepartmentRepository) Create(ctx context.Context, department models.Department) (int64, error) {
	departmentData := []interface{}{
		department.Name,
		department.Phone,
	}

	var depId int64
	err := d.db.GetContext(ctx, &depId, depqueries.CreateDepartment, departmentData...)
	if err != nil {
		return 0, err
	}

	return depId, nil
}

func (d *DepartmentRepository) Delete(ctx context.Context, departmentId int64) error {
	_, err := d.db.ExecContext(ctx, depqueries.DeleteDepartment, departmentId)
	if err != nil {
		return err
	}

	return nil
}
