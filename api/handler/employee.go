package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"test-task-sw/entity"
	"test-task-sw/lib/thttp"
	"test-task-sw/service"
)

type createEmployeeRequest struct {
	Id              int    `json:"id" binding:"required" example:"1"`
	Name            string `json:"name" binding:"required" example:"Ivan"`
	Surname         string `json:"surname" binding:"required" example:"Ivanov"`
	Phone           string `json:"phone" binding:"required" example:"+7(987)6667788"`
	CompanyId       int    `json:"company_id" binding:"required" example:"1"`
	PassportType    string `json:"passport_type" binding:"required" example:"Russian passport"`
	PassportNumber  string `json:"passport_number" binding:"required" example:"1122 112233"`
	DepartmentName  string `json:"department_name" binding:"required" example:"First Department"`
	DepartmentPhone string `json:"department_phone" binding:"required" example:"+7(987)1112233"`
}

// CreateEmployee godoc
// @Summary     Создание работника
// @Tags		Employee
// @Accept      json
// @Produce     json
// @Param       request body createEmployeeRequest true "Данные пользователя"
// @Success     200
// @Router      /api/employee/create [post]
func CreateEmployee(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createEmployeeRequest
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		employee := entity.Employee{
			Id:        req.Id,
			Name:      req.Name,
			Surname:   req.Surname,
			Phone:     req.Phone,
			CompanyId: req.CompanyId,
			Passport: entity.Passport{
				Type:   req.PassportType,
				Number: req.PassportNumber,
			},
			Department: entity.Department{
				Name:  req.DepartmentName,
				Phone: req.DepartmentPhone,
			},
		}

		id, err := employeeService.Create(c, employee)
		switch {
		case err == nil:
		case errors.Is(err, service.ErrAlreadyExists):
			thttp.ErrorResponse(c, http.StatusUnauthorized, "Пользователь уже существует")
			return
		default:
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}
		thttp.OkResponseWithResult(c, id)
		return
	}
}

type deleteEmployeeUri struct {
	EmployeeId string `uri:"employeeId"`
}

// DeleteEmployee godoc
// @Summary     Удаление работника
// @Tags		Employee
// @Accept      json
// @Produce     json
// @Param       employeeId path string true "Идентификатор работника"
// @Success     200 {object} thttp.ResponseOk
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Router      /api/employee/delete/{employeeId} [delete]
func DeleteEmployee(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req deleteEmployeeUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		id, err := strconv.Atoi(req.EmployeeId)
		if err != nil {
			return
		}

		err = employeeService.DeleteEmployee(c, int64(id))
		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusNotFound, service.ErrNotFound.Error())
			return
		default:
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		thttp.OkResponse(c)
	}
}

type getListEmployeesByCompanyIdUri struct {
	CompanyId string `uri:"companyId"`
}

type getListEmployeesByCompanyIdResponse []getListEmployeesByCompanyIdElement

type getListEmployeesByCompanyIdElement struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Phone           string `json:"phone"`
	CompanyId       int    `json:"company_id"`
	PassportType    string `json:"passport_type"`
	PassportNumber  string `json:"passport_number"`
	DepartmentName  string `json:"department_name"`
	DepartmentPhone string `json:"department_phone"`
}

func newListEmployeesByCompanyIdResponse(employees []entity.Employee) getListEmployeesByCompanyIdResponse {
	var response = make(getListEmployeesByCompanyIdResponse, 0, len(employees))
	for _, employee := range employees {
		response = append(response, getListEmployeesByCompanyIdElement{
			Id:              employee.Id,
			Name:            employee.Name,
			Surname:         employee.Surname,
			Phone:           employee.Phone,
			CompanyId:       employee.CompanyId,
			PassportType:    employee.Passport.Type,
			PassportNumber:  employee.Passport.Number,
			DepartmentName:  employee.Department.Name,
			DepartmentPhone: employee.Department.Phone,
		})
	}
	return response
}

// ListEmployeesByCompanyId godoc
// @Summary     Получение работников по id компании
// @Tags		Employee
// @Accept      json
// @Produce     json
// @Param       companyId path string true "Идентификатор компании"
// @Success     200 {object} thttp.ResponseWithDetails[getListEmployeesByCompanyIdResponse]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Router      /api/employee/list/{companyId} [get]
func ListEmployeesByCompanyId(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getListEmployeesByCompanyIdUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		companyId, _ := strconv.Atoi(req.CompanyId)

		employees, err := employeeService.GetEmployeeListByCompanyId(c, companyId)
		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusNotFound, service.ErrNotFound.Error())
			return
		default:
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusInternalServerError, "internal  server error")
			return
		}

		response := newListEmployeesByCompanyIdResponse(employees)
		thttp.OkResponseWithResult(c, response)
	}
}

type listEmployeesByDepartmentUri struct {
	DepName string `uri:"depName"`
}

type listEmployeesByDepartmentResponse []listEmployeesByDepartmentElement

type listEmployeesByDepartmentElement struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Phone           string `json:"phone"`
	CompanyId       int    `json:"company_id"`
	PassportType    string `json:"passport_type"`
	PassportNumber  string `json:"passport_number"`
	DepartmentName  string `json:"department_name"`
	DepartmentPhone string `json:"department_phone"`
}

func newListEmployeesByDepartmentResponse(employees []entity.Employee) listEmployeesByDepartmentResponse {
	var response = make(listEmployeesByDepartmentResponse, 0, len(employees))
	for _, employee := range employees {
		response = append(response, listEmployeesByDepartmentElement{
			Id:              employee.Id,
			Name:            employee.Name,
			Surname:         employee.Surname,
			Phone:           employee.Phone,
			CompanyId:       employee.CompanyId,
			PassportType:    employee.Passport.Type,
			PassportNumber:  employee.Passport.Number,
			DepartmentName:  employee.Department.Name,
			DepartmentPhone: employee.Department.Phone,
		})
	}
	return response
}

// ListEmployeesByDepartment godoc
// @Summary     Получение работников по отделу
// @Tags		Employees
// @Accept      json
// @Produce     json
// @Param       depName path string true "Название отдела"
// @Success     200 {object} thttp.ResponseWithDetails[listEmployeesByDepartmentResponse]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Router      /api/employee/list/department/{depName} [get]
func ListEmployeesByDepartment(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listEmployeesByDepartmentUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		employees, err := employeeService.GetEmployeeListByDepartmentName(c, req.DepName)

		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusNotFound, service.ErrNotFound.Error())
			return
		default:
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusInternalServerError, "internal  server error")
			return
		}

		response := newListEmployeesByDepartmentResponse(employees)
		thttp.OkResponseWithResult(c, response)
	}
}

type updateEmployeeUri struct {
	EmployeeId string `uri:"employeeId"`
}

type updateEmployeeRequest struct {
	Name            string `json:"name,omitempty"`
	Surname         string `json:"surname,omitempty"`
	Phone           string `json:"phone,omitempty"`
	CompanyId       *int   `json:"company_id,omitempty"`
	PassportType    string `json:"passport_type,omitempty"`
	PassportNumber  string `json:"passport_number,omitempty"`
	DepartmentName  string `json:"department_name,omitempty"`
	DepartmentPhone string `json:"department_phone,omitempty"`
}

type updateEmployeeResponse struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Phone           string `json:"phone"`
	CompanyId       int    `json:"company_id"`
	PassportType    string `json:"passport_type"`
	PassportNumber  string `json:"passport_number"`
	DepartmentName  string `json:"department_name"`
	DepartmentPhone string `json:"department_phone"`
}

func newUpdateEmployeeResponse(employee entity.Employee) updateEmployeeResponse {
	return updateEmployeeResponse{
		Id:              employee.Id,
		Name:            employee.Name,
		Surname:         employee.Surname,
		Phone:           employee.Phone,
		CompanyId:       employee.CompanyId,
		PassportType:    employee.Passport.Type,
		PassportNumber:  employee.Passport.Number,
		DepartmentName:  employee.Department.Name,
		DepartmentPhone: employee.Department.Phone,
	}
}

// UpdateEmployee godoc
// @Summary     Изменение данных работника
// @Tags		Employee
// @Accept      json
// @Produce     json
// @Param       employeeId path string true "Идентификатор работника"
// @Param 		request body updateEmployeeRequest true "da"
// @Success     200 {object} thttp.ResponseWithDetails[updateEmployeeResponse]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     409 {object} thttp.ResponseError "Already exists"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Router      /api/employee/update/{employeeId} [put]
func UpdateEmployee(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id updateEmployeeUri
		if err := c.ShouldBindUri(&id); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		var req updateEmployeeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		employeeId, err := strconv.ParseInt(id.EmployeeId, 10, 64)
		if err != nil {
			return
		}

		employee := entity.Employee{
			Name:      req.Name,
			Surname:   req.Surname,
			Phone:     req.Phone,
			CompanyId: *req.CompanyId,
			Passport: entity.Passport{
				Type:   req.PassportType,
				Number: req.PassportNumber,
			},
			Department: entity.Department{
				Name:  req.DepartmentName,
				Phone: req.DepartmentPhone,
			},
		}

		err = employeeService.UpdateEmployee(c, employeeId, employee)
		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusNotFound, service.ErrNotFound.Error())
			return
		default:
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		//response :=
		thttp.OkResponse(c)
	}
}
