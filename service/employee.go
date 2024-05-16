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

func (u *EmployeeService) Create(ctx context.Context, employee entity.Employee) (int64, error) {
	//isUserExists, err := u.isUserExists(ctx, employee.Id)
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
	passId, err := u.passportService.CreatePassport(ctx, passport)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	//userId, err := u.employeeRepo.Create(ctx, employee)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	return passId, nil
}

func (u *EmployeeService) isUserExists(ctx context.Context, userId int) (bool, error) {

	_, err := u.employeeRepo.GetEmployee(ctx, userId)

	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return false, nil
	default:
		return false, err
	}

	return true, nil
}
