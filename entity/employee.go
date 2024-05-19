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

func (e *Employee) PartialUpdate(updateModel Employee) {
	if len(updateModel.Name) != 0 {
		e.Name = updateModel.Name
	}
	if len(updateModel.Surname) != 0 {
		e.Surname = updateModel.Surname
	}
	if len(updateModel.Phone) != 0 {
		e.Phone = updateModel.Phone
	}
	if updateModel.CompanyId != 0 {
		e.CompanyId = updateModel.CompanyId
	}

	e.Passport.PartialUpdate(updateModel.Passport)
	e.Department.PartialUpdate(updateModel.Department)
}
