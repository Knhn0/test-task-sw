package passportqueries

import _ "embed"

//go:embed insert_passport_data.sql
var InsertPassportData string

//go:embed delete_passport.sql
var DeletePassport string

//go:embed update_passport.sql
var UpdatePassport string
