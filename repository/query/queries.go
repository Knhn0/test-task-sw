package query

import _ "embed"

//go:embed insert_employee_data.sql
var InsertEmployeeData string

//go:embed insert_passport_data.sql
var InsertPassportData string

//go:embed insert_department_data.sql
var InsertDepartmentData string

//go:embed delete_passport.sql
var DeletePassport string

//go:embed delete_department.sql
var DeleteDepartment string

//go:embed delete_employee.sql
var DeleteEmployee string

//go:embed get_employee.sql
var GetEmployee string
