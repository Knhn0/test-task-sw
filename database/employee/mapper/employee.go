package mapper

import (
	"test-task-sw/database/employee/models"
	smodels "test-task-sw/service/models"
)

func MapEmployee(employee models.Employee) smodels.Employee {
	return smodels.Employee{
		Id:        employee.Id,
		Name:      employee.Name,
		Surname:   employee.Surname,
		Phone:     employee.Phone,
		CompanyId: employee.CompanyId,
		Passport: smodels.Passport{
			Type:   employee.PassportType,
			Number: employee.PassportNumber,
		},
		Department: smodels.Department{
			Name:  employee.DepartmentName,
			Phone: employee.DepartmentPhone,
		},
	}
}

func EmployeesSlice(employeeModels []models.Employee) []smodels.Employee {
	var employees = make([]smodels.Employee, 0, len(employeeModels))
	for _, employee := range employeeModels {
		employees = append(employees, MapEmployee(employee))
	}

	return employees
}
