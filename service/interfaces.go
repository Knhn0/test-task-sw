package service

import (
	"context"
	"test-task-sw/service/models"
)

type employeeRepo interface {
	Create(ctx context.Context, employee models.Employee) (int64, error)
	Delete(ctx context.Context, employeeId int64) error
	GetListByCompanyId(ctx context.Context, companyId int) ([]models.Employee, error)
	GetListByDepartmentName(ctx context.Context, depName string) ([]models.Employee, error)
	UpdateEmployee(ctx context.Context, employeeId int64, employee models.Employee) error
	Get(ctx context.Context, employeeId int64) (models.Employee, error)
}

type passportRepo interface {
	Create(ctx context.Context, passport models.Passport) (int64, error)
	Delete(ctx context.Context, passportId int64) error
}

type departmentRepo interface {
	Create(ctx context.Context, department models.Department) (int64, error)
	Delete(ctx context.Context, departmentId int64) error
}
