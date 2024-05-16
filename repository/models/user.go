package models

import "github.com/google/uuid"

type User struct {
	Id           int       `db:"id"`
	Name         string    `db:"user_name"`
	Surname      string    `db:"surname"`
	Phone        string    `db:"phone"`
	CompanyId    int       `db:"company_id"`
	PassportId   uuid.UUID `db:"passport_id"`
	DepartmentId uuid.UUID `db:"department_id"`
}

type Passport struct {
}
