package models

type Department struct {
	Name  string `json:"department_name"`
	Phone string `json:"department_phone"`
}

func (d *Department) PartialUpdate(updateModel Department) {
	if len(updateModel.Name) != 0 {
		d.Name = updateModel.Name
	}
	if len(updateModel.Phone) != 0 {
		d.Phone = updateModel.Phone
	}
}
