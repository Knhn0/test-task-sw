package queries

import _ "embed"

//go:embed create_employee.sql
var CreateEmployee string

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
