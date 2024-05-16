package service

type DepartmentService struct {
	departmentRepo departmentRepo
}

func NewDepartmentService(repo departmentRepo) *DepartmentService {
	return &DepartmentService{
		repo,
	}
}
