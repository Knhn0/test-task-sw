package models

type Passport struct {
	Type   string `json:"passport_type"`
	Number string `json:"passport_number"`
}

func (p *Passport) PartialUpdate(updateModel Passport) {
	if len(updateModel.Type) != 0 {
		p.Type = updateModel.Type
	}
	if len(updateModel.Number) != 0 {
		p.Number = updateModel.Number
	}
}
