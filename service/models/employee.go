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
	return _passportRegex.Match([]byte(e.Passport.Number)) &&
		_phoneRegex.Match([]byte(e.Phone)) &&
		_phoneRegex.Match([]byte(e.Department.Phone))

}

func (e *Employee) PartialUpdate(updateModel Employee) {
	if len(updateModel.Name) != 0 {
		e.Name = updateModel.Name
	}
	if len(updateModel.Surname) != 0 {
		e.Surname = updateModel.Surname
	}
	if len(updateModel.Phone) != 0 {
		e.Phone = updateModel.Phone
	}
	if updateModel.CompanyId != 0 {
		e.CompanyId = updateModel.CompanyId
	}

	e.Passport.PartialUpdate(updateModel.Passport)
	e.Department.PartialUpdate(updateModel.Department)
}
