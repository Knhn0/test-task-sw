package repository

import "test-task-sw/lib/tpostgres"

type DepartmentRepository struct {
	db *tpostgres.Postgres
}

func NewDepartmentRepository(db *tpostgres.Postgres) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}
