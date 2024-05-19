package depqueries

import _ "embed"

//go:embed create_department.sql
var CreateDepartment string

//go:embed delete_department.sql
var DeleteDepartment string

//go:embed update_department.sql
var UpdateDepartment string
