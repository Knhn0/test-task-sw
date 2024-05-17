package service

import (
	"context"
	"database/sql"
	"errors"
	"test-task-sw/entity"
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

func (e *EmployeeService) Create(ctx context.Context, employee entity.Employee) (int64, error) {
	//isEmployeeExists, err := e.isEmployeeExists(ctx, int64(employee.Id))
	//if err != nil {
	//	return 0, err
	//}
	//
	//if isEmployeeExists {
	//	err = ErrAlreadyExists
	//	return 0, err
	//}

	passport := entity.Passport{
		Type:   employee.Passport.Type,
		Number: employee.Passport.Number,
	}
	passId, err := e.passportService.CreatePassport(ctx, passport)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	department := entity.Department{
		Name:  employee.Department.Name,
		Phone: employee.Department.Phone,
	}
	depId, err := e.departmentService.CreateDepartment(ctx, department)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	employeeId, err := e.employeeRepo.Create(ctx, employee, passId, depId)
	switch {
	case err == nil:
	default:
		e.passportService.DeletePassport(ctx, passId)
		return 0, err
	}

	return employeeId, nil
}

func (e *EmployeeService) GetEmployee(ctx context.Context, employeeId int64) (entity.EmployeeTransfer, error) {
	employee, err := e.employeeRepo.GetEmployee(ctx, employeeId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return entity.EmployeeTransfer{}, ErrNotFound
	default:
		return entity.EmployeeTransfer{}, err
	}
	return employee, nil
}

func (e *EmployeeService) DeleteEmployee(ctx context.Context, employeeId int64) error {
	err := e.employeeRepo.Delete(ctx, employeeId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
	return nil
}

func (e *EmployeeService) isEmployeeExists(ctx context.Context, employeeId int64) (bool, error) {

	_, err := e.employeeRepo.GetEmployee(ctx, employeeId)

	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return false, nil
	default:
		return false, err
	}

	return true, nil
}
