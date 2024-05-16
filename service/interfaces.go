package service

import (
	"context"
	"test-task-sw/entity"
)

type employeeRepo interface {
	GetEmployee(ctx context.Context, employeeId int) (entity.Employee, error)
	Create(ctx context.Context, employee entity.Employee) (int64, error)
}

type passportRepo interface {
	Create(ctx context.Context, passport entity.Passport) (int64, error)
}

type departmentRepo interface {
}
