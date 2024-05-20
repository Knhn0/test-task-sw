package passportqueries

import _ "embed"

//go:embed create_passport.sql
var CreatePassport string

//go:embed delete_passport.sql
var DeletePassport string

//go:embed update_passport.sql
var UpdatePassport string
