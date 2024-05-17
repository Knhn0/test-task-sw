package mapper

import (
	"test-task-sw/entity"
	"test-task-sw/repository/models"
)

func EmployeeTransferMapFromDb(employee models.Employee) entity.EmployeeTransfer {
	return entity.EmployeeTransfer{
		Id:           employee.Id,
		Name:         employee.Name,
		Surname:      employee.Surname,
		Phone:        employee.Phone,
		CompanyId:    employee.CompanyId,
		PassportId:   employee.PassportId,
		DepartmentId: employee.DepartmentId,
	}
}

func MapEmployeeForList(employee models.EmployeeForList) entity.Employee {
	return entity.Employee{
		Id:        employee.Id,
		Name:      employee.Name,
		Surname:   employee.Surname,
		Phone:     employee.Phone,
		CompanyId: employee.CompanyId,
		Passport: entity.Passport{
			Type:   employee.PassportType,
			Number: employee.PassportNumber,
		},
		Department: entity.Department{
			Name:  employee.DepartmentName,
			Phone: employee.DepartmentPhone,
		},
	}
}

func EmployeesSlice(employeeModels []models.EmployeeForList) []entity.Employee {
	var employees = make([]entity.Employee, 0, len(employeeModels))
	for _, employee := range employeeModels {
		employees = append(employees, MapEmployeeForList(employee))
	}
	return employees
}
