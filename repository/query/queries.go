package query

import _ "embed"

//go:embed insert_employee_data.sql
var InsertEmployeeData string

//go:embed insert_passport_data.sql
var InsertPassportData string

//go:embed insert_department_data.sql
var InsertDepartmentData string
