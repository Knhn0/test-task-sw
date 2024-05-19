package entity

type Passport struct {
	Type   string
	Number string
}

func (p *Passport) PartialUpdate(updateModel Passport) {
	if len(updateModel.Type) != 0 {
		p.Type = updateModel.Type
	}
	if len(updateModel.Number) != 0 {
		p.Number = updateModel.Number
	}
}
