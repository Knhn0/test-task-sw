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
	//isUserExists, err := e.isUserExists(ctx, employee.Id)
	//if err != nil {
	//	return 0, err
	//}
	//
	//if isUserExists {
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
		return 0, err
	}

	return employeeId, nil
}

func (e *EmployeeService) isUserExists(ctx context.Context, userId int) (bool, error) {

	_, err := e.employeeRepo.GetEmployee(ctx, userId)

	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return false, nil
	default:
		return false, err
	}

	return true, nil
}
