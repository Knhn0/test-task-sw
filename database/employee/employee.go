package employee

import (
	"context"
	depqueries "test-task-sw/database/department/queries"
	"test-task-sw/database/employee/mapper"
	"test-task-sw/database/employee/models"
	"test-task-sw/database/employee/queries"
	passqueries "test-task-sw/database/passport/queries"
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

func (e *EmployeeRepository) Create(ctx context.Context, employee entity.Employee) (int64, error) {

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var passId int64
	err = tx.QueryRowContext(ctx, passqueries.InsertPassportData, employee.Passport.Type, employee.Passport.Number).Scan(&passId)
	if err != nil {
		return 0, err
	}

	var depId int64
	err = tx.QueryRowContext(ctx, depqueries.InsertDepartmentData, employee.Department.Name, employee.Department.Phone).Scan(&depId)
	if err != nil {
		return 0, err
	}

	var employeeId int64
	err = tx.QueryRowContext(ctx, queries.InsertEmployeeData, employee.Name, employee.Surname, employee.Phone, employee.CompanyId, passId, depId).Scan(&employeeId)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return employeeId, nil
}

func (e *EmployeeRepository) Delete(ctx context.Context, employeeId int64) error {

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var passportId int64
	err = tx.QueryRowContext(ctx, queries.DeleteEmployee, employeeId).Scan(&passportId)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, passqueries.DeletePassport, passportId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepository) GetListByCompanyId(ctx context.Context, companyId int) ([]entity.Employee, error) {
	var employeeModel []models.Employee
	if err := e.db.SelectContext(ctx, &employeeModel, queries.GetListByCompanyId, companyId); err != nil {
		return nil, err
	}
	return mapper.EmployeesSlice(employeeModel), nil
}

func (e *EmployeeRepository) GetListByDepartmentName(ctx context.Context, depName string) ([]entity.Employee, error) {
	var employeeModel []models.Employee
	if err := e.db.SelectContext(ctx, &employeeModel, queries.GetListEmployeesByDepName, depName); err != nil {
		return nil, err
	}
	return mapper.EmployeesSlice(employeeModel), nil
}

func (e *EmployeeRepository) Get(ctx context.Context, employeeId int64) (entity.Employee, error) {
	var employeeModel models.Employee
	if err := e.db.GetContext(ctx, &employeeModel, queries.GetEmployee, employeeId); err != nil {
		return entity.Employee{}, err
	}
	return mapper.MapEmployee(employeeModel), nil
}

func (e *EmployeeRepository) UpdateEmployee(ctx context.Context, employeeId int64, employee entity.Employee) error {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var passId int
	var depId int

	err = tx.QueryRowContext(ctx, queries.UpdateEmployee, employeeId, employee.Name, employee.Surname, employee.Phone, employee.CompanyId).Scan(&passId, &depId)
	if err != nil {
		return err
	}

	argsForPassport := []interface{}{
		passId,
		employee.Passport.Type,
		employee.Passport.Number,
	}
	_, err = tx.ExecContext(ctx, passqueries.UpdatePassport, argsForPassport...)
	if err != nil {
		return err
	}

	argsForDepartment := []interface{}{
		depId,
		employee.Department.Name,
		employee.Department.Phone,
	}
	_, err = tx.ExecContext(ctx, depqueries.UpdateDepartment, argsForDepartment...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
