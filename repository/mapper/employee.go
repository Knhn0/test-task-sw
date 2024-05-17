package mapper

import (
	"test-task-sw/entity"
	"test-task-sw/repository/models"
)

func EmployeeMapFromDb(employee models.Employee) entity.EmployeeTransfer {
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
