package entity

type Employee struct {
	Id         int
	Name       string
	Surname    string
	Phone      string
	CompanyId  int
	Passport   Passport
	Department Department
}

type EmployeeTransfer struct {
	Id           int
	Name         string
	Surname      string
	Phone        string
	CompanyId    int
	PassportId   int
	DepartmentId int
}

type EmployeeForList struct {
	Id              int
	Name            string
	Surname         string
	Phone           string
	CompanyId       int
	PassportType    string
	PassportNumber  string
	DepartmentName  string
	DepartmentPhone string
}
