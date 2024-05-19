package models

type Department struct {
	Id    int    `db:"id"`
	Name  string `db:"department_name"`
	Phone string `db:"department_phone"`
}
