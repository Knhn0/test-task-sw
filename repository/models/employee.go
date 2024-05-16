package models

type Employee struct {
	Id           int    `db:"id"`
	Name         string `db:"user_name"`
	Surname      string `db:"surname"`
	Phone        string `db:"phone"`
	CompanyId    int    `db:"company_id"`
	PassportId   int    `db:"passport_id"`
	DepartmentId int    `db:"department_id"`
}
