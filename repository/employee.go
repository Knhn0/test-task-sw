package repository

import (
	"context"
	"test-task-sw/entity"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/repository/mapper"
	"test-task-sw/repository/models"
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

func (e *EmployeeRepository) GetEmployee(ctx context.Context, employeeId int64) (entity.EmployeeTransfer, error) {
	var employeeModel models.Employee
	if err := e.db.GetContext(ctx, &employeeModel, query.GetEmployee, employeeId); err != nil {
		return entity.EmployeeTransfer{}, err
	}
	return mapper.EmployeeMapFromDb(employeeModel), nil
}

func (e *EmployeeRepository) Create(ctx context.Context, employee entity.Employee, passId int64, depId int64) (int64, error) {
	args := []interface{}{
		employee.Name,
		employee.Surname,
		employee.Phone,
		employee.CompanyId,
		passId,
		depId,
	}

	var employeeId int64
	err := e.db.GetContext(ctx, &employeeId, query.InsertEmployeeData, args...)
	if err != nil {
		return 0, err
	}
	return employeeId, nil
}

func (e *EmployeeRepository) Delete(ctx context.Context, employeeId int64) error {
	employee, err := e.GetEmployee(ctx, employeeId)
	if err != nil {
		return err
	}

	_, err = e.db.ExecContext(ctx, query.DeletePassport, employee.PassportId)
	if err != nil {
		return err
	}
	return nil
}
