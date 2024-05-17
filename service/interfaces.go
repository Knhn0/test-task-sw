package service

import (
	"context"
	"test-task-sw/entity"
)

type employeeRepo interface {
	GetEmployee(ctx context.Context, employeeId int) (entity.Employee, error)
	Create(ctx context.Context, employee entity.Employee, passId int64, depId int64) (int64, error)
}

type passportRepo interface {
	Create(ctx context.Context, passport entity.Passport) (int64, error)
}

type departmentRepo interface {
	Create(ctx context.Context, department entity.Department) (int64, error)
}
