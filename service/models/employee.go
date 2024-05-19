package models

import "regexp"

var (
	_passportRegex = regexp.MustCompile(`^[0-9]{10,}$`)
	_phoneRegex    = regexp.MustCompile(`^[0-9()+]{14,}$`)
)

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Phone      string `json:"phone"`
	CompanyId  int    `json:"company_id"`
	Passport   Passport
	Department Department
}

func (e *Employee) Validate() bool {
	if !_passportRegex.Match([]byte(e.Passport.Number)) {
		return false
	}
	if !_phoneRegex.Match([]byte(e.Phone)) {
		return false
	}
	if !_phoneRegex.Match([]byte(e.Department.Phone)) {
		return false
	}

	return true
}
