package models

type Employee struct {
	Id           int    `db:"id"`
	Name         string `db:"employee_name"`
	Surname      string `db:"surname"`
	Phone        string `db:"phone"`
	CompanyId    int    `db:"company_id"`
	PassportId   int    `db:"passport_id"`
	DepartmentId int    `db:"department_id"`
}

type EmployeeForList struct {
	Id              int    `db:"id"`
	Name            string `db:"employee_name"`
	Surname         string `db:"surname"`
	Phone           string `db:"phone"`
	CompanyId       int    `db:"company_id"`
	PassportId      int    `db:"passport_id"`
	PassportType    string `db:"passport_type"`
	PassportNumber  string `db:"passport_number"`
	DepartmentId    int    `db:"department_id"`
	DepartmentName  string `db:"department_name"`
	DepartmentPhone string `db:"department_phone"`
}
