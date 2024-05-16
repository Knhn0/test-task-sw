package repository

import (
	"context"
	"test-task-sw/entity"
	"test-task-sw/lib/tpostgres"
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
func (u *EmployeeRepository) Create(ctx context.Context, employee entity.Employee) (int64, error) {

	return 1, nil
}
