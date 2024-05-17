package repository

import (
	"context"
	"test-task-sw/entity"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/repository/query"
)

type EmployeeRepository struct {
	db *tpostgres.Postgres
}

func NewUserRepository(db *tpostgres.Postgres) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (u *EmployeeRepository) GetEmployee(ctx context.Context, employeeId int) (entity.Employee, error) {
	return entity.Employee{}, nil
}
func (u *EmployeeRepository) Create(ctx context.Context, employee entity.Employee, passId int64, depId int64) (int64, error) {
	args := []interface{}{
		employee.Name,
		employee.Surname,
		employee.Phone,
		employee.CompanyId,
		passId,
		depId,
	}

	_, err := u.db.ExecContext(ctx, query.InsertEmployeeData, args...)
	if err != nil {
		return 0, err
	}

	var employeeId int64
	err = u.db.GetContext(ctx, &employeeId, query.InsertEmployeeData, args...)

	return employeeId, nil
}
