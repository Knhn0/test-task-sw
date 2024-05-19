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

//go:embed get_list_by_company_id.sql
var GetListByCompanyId string

//go:embed get_list_employees_by_dep_name.sql
var GetListEmployeesByDepName string

//go:embed update_employee.sql
var UpdateEmployee string

//go:embed update_passport.sql
var UpdatePassport string

//go:embed update_department.sql
var UpdateDepartment string
