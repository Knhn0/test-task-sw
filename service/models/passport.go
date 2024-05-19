package models

type Passport struct {
	Type   string `json:"passport_type"`
	Number string `json:"passport_number"`
}
