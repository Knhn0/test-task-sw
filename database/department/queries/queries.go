package depqueries

import _ "embed"

//go:embed insert_department_data.sql
var InsertDepartmentData string

//go:embed delete_department.sql
var DeleteDepartment string

//go:embed update_department.sql
var UpdateDepartment string
