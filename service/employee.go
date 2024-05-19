package service

import (
	"context"
	"database/sql"
	"errors"
	"test-task-sw/service/models"
)

type EmployeeService struct {
	passportService   *PassportService
	departmentService *DepartmentService
	employeeRepo      employeeRepo
}

func NewUserService(repo employeeRepo, depServ *DepartmentService, passServ *PassportService) *EmployeeService {
	return &EmployeeService{
		passServ,
		depServ,
		repo,
	}
}

func (e *EmployeeService) Create(ctx context.Context, employee models.Employee) (int64, error) {

	employeeId, err := e.employeeRepo.Create(ctx, employee)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	return employeeId, nil
}

func (e *EmployeeService) GetEmployee(ctx context.Context, employeeId int64) (models.Employee, error) {
	employee, err := e.employeeRepo.Get(ctx, employeeId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return models.Employee{}, ErrNotFound
	default:
		return models.Employee{}, err
	}

	return employee, nil
}

func (e *EmployeeService) DeleteEmployee(ctx context.Context, employeeId int64) error {
	isEmployeeExists, err := e.isEmployeeExists(ctx, employeeId)
	if err != nil {
		return err
	}

	if !isEmployeeExists {
		err = ErrNotFound
		return err
	}

	err = e.employeeRepo.Delete(ctx, employeeId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}

	return nil
}

func (e *EmployeeService) GetEmployeeListByCompanyId(ctx context.Context, companyId int) ([]models.Employee, error) {
	// нужно добавить ifExists
	employees, err := e.employeeRepo.GetListByCompanyId(ctx, companyId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return []models.Employee{}, ErrNotFound
	default:
		return []models.Employee{}, err
	}

	return employees, nil
}

func (e *EmployeeService) GetEmployeeListByDepartmentName(ctx context.Context, departmentName string) ([]models.Employee, error) {
	employees, err := e.employeeRepo.GetListByDepartmentName(ctx, departmentName)
	if len(employees) == 0 {
		return []models.Employee{}, ErrNotFound
	}
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return []models.Employee{}, ErrNotFound
	default:
		return []models.Employee{}, err
	}

	return employees, nil
}

func (e *EmployeeService) isEmployeeExists(ctx context.Context, employeeId int64) (bool, error) {

	_, err := e.employeeRepo.Get(ctx, employeeId)

	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return false, nil
	default:
		return false, err
	}

	return true, nil
}

func (e *EmployeeService) UpdateEmployee(ctx context.Context, employeeId int64, updEmployee models.Employee) error {

	lastDbModel, err := e.employeeRepo.Get(ctx, employeeId)
	if err != nil {
		return err
	}

	lastDbModel.PartialUpdate(updEmployee)
	err = e.employeeRepo.UpdateEmployee(ctx, employeeId, lastDbModel)

	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return nil
	default:
		return err
	}

	return nil
}
