package service

import (
	"context"
	"test-task-sw/entity"
)

type employeeRepo interface {
	//GetEmployee(ctx context.Context, employeeId int64) (entity.EmployeeTransfer, error)
	Create(ctx context.Context, employee entity.Employee) (int64, error)
	Delete(ctx context.Context, employeeId int64) error
	GetListByCompanyId(ctx context.Context, companyId int) ([]entity.Employee, error)
	GetListByDepartmentName(ctx context.Context, depName string) ([]entity.Employee, error)
	UpdateEmployee(ctx context.Context, employeeId int64, employee entity.Employee) error
	Get(ctx context.Context, employeeId int64) (entity.Employee, error)
}

type passportRepo interface {
	Create(ctx context.Context, passport entity.Passport) (int64, error)
	Delete(ctx context.Context, passportId int64) error
}

type departmentRepo interface {
	Create(ctx context.Context, department entity.Department) (int64, error)
	Delete(ctx context.Context, departmentId int64) error
}
