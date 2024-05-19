package models

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Phone      string `json:"phone"`
	CompanyId  int    `json:"company_id"`
	Passport   Passport
	Department Department
}
