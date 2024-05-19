package models

type Passport struct {
	Id     int    `db:"id"`
	Type   string `db:"passport_type"`
	Number string `db:"passport_number"`
}
