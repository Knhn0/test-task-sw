package service

import (
	"context"
	"test-task-sw/entity"
)

type DepartmentService struct {
	departmentRepo departmentRepo
}

func NewDepartmentService(repo departmentRepo) *DepartmentService {
	return &DepartmentService{
		repo,
	}
}

func (d *DepartmentService) CreateDepartment(ctx context.Context, department entity.Department) (int64, error) {
	depId, err := d.departmentRepo.Create(ctx, department)
	switch {
	case err == nil:
	default:
		return 0, err
	}

	return depId, nil
}
