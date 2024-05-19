package employee

import (
	"context"
	"test-task-sw/database/department"
	"test-task-sw/database/employee/mapper"
	"test-task-sw/database/employee/models"
	"test-task-sw/database/employee/queries"
	"test-task-sw/database/passport"
	"test-task-sw/lib/tpostgres"
	smodels "test-task-sw/service/models"
)

type Repo interface {
	Create(ctx context.Context, employee smodels.Employee) (int64, error)
	Delete(ctx context.Context, employeeId int64) error
	GetListByCompanyId(ctx context.Context, companyId int) ([]smodels.Employee, error)
	GetListByDepartmentName(ctx context.Context, depName string) ([]smodels.Employee, error)
	Update(ctx context.Context, employeeId int64, employee smodels.Employee) error
	Get(ctx context.Context, employeeId int64) (smodels.Employee, error)
}

type Impl struct {
	db             *tpostgres.Postgres
	passportRepo   *passport.Impl
	departmentRepo *department.Impl
}

func NewRepository(db *tpostgres.Postgres, passportRepo *passport.Impl, departmentRepo *department.Impl) Repo {
	return &Impl{
		db:             db,
		passportRepo:   passportRepo,
		departmentRepo: departmentRepo,
	}
}

func (e *Impl) Create(ctx context.Context, employee smodels.Employee) (int64, error) {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var employeeId int64

	passId, err := e.passportRepo.Create(ctx, tx, employee.Passport)
	if err != nil {
		return 0, err
	}

	depId, err := e.departmentRepo.Create(ctx, tx, employee.Department)
	if err != nil {
		return 0, err
	}

	err = tx.QueryRowContext(ctx, queries.CreateEmployee, employee.Name, employee.Surname, employee.Phone, employee.CompanyId, passId, depId).Scan(&employeeId)
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return employeeId, nil
}

func (e *Impl) Delete(ctx context.Context, employeeId int64) error {
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

	err = e.passportRepo.Delete(ctx, tx, passportId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (e *Impl) GetListByCompanyId(ctx context.Context, companyId int) ([]smodels.Employee, error) {
	var employeeModel []models.Employee
	if err := e.db.SelectContext(ctx, &employeeModel, queries.GetListByCompanyId, companyId); err != nil {
		return nil, err
	}

	return mapper.EmployeesSlice(employeeModel), nil
}

func (e *Impl) GetListByDepartmentName(ctx context.Context, depName string) ([]smodels.Employee, error) {
	var employeeModel []models.Employee
	if err := e.db.SelectContext(ctx, &employeeModel, queries.GetListEmployeesByDepName, depName); err != nil {
		return nil, err
	}

	return mapper.EmployeesSlice(employeeModel), nil
}

func (e *Impl) Get(ctx context.Context, employeeId int64) (smodels.Employee, error) {
	var employeeModel models.Employee
	if err := e.db.GetContext(ctx, &employeeModel, queries.GetEmployee, employeeId); err != nil {
		return smodels.Employee{}, err
	}

	return mapper.MapEmployee(employeeModel), nil
}

func (e *Impl) Update(ctx context.Context, employeeId int64, employee smodels.Employee) error {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var passportId int64
	var departmentId int64

	err = tx.QueryRowContext(ctx, queries.UpdateEmployee, employeeId, employee.Name, employee.Surname, employee.Phone, employee.CompanyId).Scan(&passportId, &departmentId)
	if err != nil {
		return err
	}

	err = e.passportRepo.Update(ctx, tx, passportId, employee.Passport)
	if err != nil {
		return err
	}

	err = e.departmentRepo.Update(ctx, tx, departmentId, employee.Department)
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
