package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"test-task-sw/lib/thttp"
	"test-task-sw/service"
	"test-task-sw/service/models"
)

// CreateEmployee godoc
// @Summary     Создание работника
// @Tags		Employee
// @Accept      json
// @Produce     json
// @Param       input body models.Employee true "Данные пользователя"
// @Success     200
// @Router      /api/employee/create [post]
func CreateEmployee(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Employee
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			logger.Error(err.Error())
			thttp.SendErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		if !req.Validate() {
			thttp.SendErrorResponse(c, http.StatusBadRequest, "validation error")
			return
		}

		id, err := employeeService.Create(c, req)
		switch {
		case err == nil:
		case errors.Is(err, service.ErrAlreadyExists):
			thttp.SendErrorResponse(c, http.StatusUnauthorized, "employee not found")
			return
		default:
			logger.Error(err.Error())
			thttp.SendErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		thttp.SendOkResponse(c, id)
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
			thttp.SendErrorResponse(c, http.StatusBadRequest, "bad body")
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
			thttp.SendErrorResponse(c, http.StatusNotFound, "employee not found")
			return
		default:
			logger.Error(err.Error())
			thttp.SendErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}

		thttp.SendOkResponse(c, nil)
	}
}

type getListEmployeesByCompanyIdUri struct {
	CompanyId string `uri:"companyId"`
}

// ListEmployeesByCompanyId godoc
// @Summary     Получение работников по id компании
// @Tags		Employee
// @Accept      json
// @Produce     json
// @Param       companyId path string true "Идентификатор компании"
// @Success     200 {object} thttp.ResponseWithDetails[[]models.Employee]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Router      /api/employee/list/company/{companyId} [get]
func ListEmployeesByCompanyId(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getListEmployeesByCompanyIdUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.SendErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		companyId, _ := strconv.Atoi(req.CompanyId)

		employees, err := employeeService.GetEmployeeListByCompanyId(c, companyId)
		switch {
		case err == nil:
		default:
			logger.Error(err.Error())
			thttp.SendErrorResponse(c, http.StatusInternalServerError, "internal  server error")
			return
		}

		thttp.SendOkResponse(c, employees)
	}
}

type listEmployeesByDepartmentUri struct {
	DepName string `uri:"depName"`
}

// ListEmployeesByDepartment godoc
// @Summary     Получение работников по отделу
// @Tags		Employees
// @Accept      json
// @Produce     json
// @Param       depName path string true "Название отдела"
// @Success     200 {object} thttp.ResponseWithDetails[[]models.Employee]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Router      /api/employee/list/department/{depName} [get]
func ListEmployeesByDepartment(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listEmployeesByDepartmentUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.SendErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		employees, err := employeeService.GetEmployeeListByDepartmentName(c, req.DepName)
		switch {
		case err == nil:
		default:
			logger.Error(err.Error())
			thttp.SendErrorResponse(c, http.StatusInternalServerError, "internal  server error")
			return
		}

		thttp.SendOkResponse(c, employees)
	}
}

type updateEmployeeUri struct {
	EmployeeId string `uri:"employeeId"`
}

// UpdateEmployee godoc
// @Summary     Изменение данных работника
// @Tags		Employee
// @Accept      json
// @Produce     json
// @Param       employeeId path string true "Идентификатор работника"
// @Param 		input body models.Employee true "Данные для обновления работника"
// @Success     200 {object} thttp.ResponseWithDetails[models.Employee]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     409 {object} thttp.ResponseError "Already exists"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Router      /api/employee/update/{employeeId} [put]
func UpdateEmployee(logger *zap.SugaredLogger, employeeService *service.EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id updateEmployeeUri
		if err := c.ShouldBindUri(&id); err != nil {
			thttp.SendErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		var req models.Employee
		if err := c.ShouldBindJSON(&req); err != nil {
			thttp.SendErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		if !req.Validate() {
			thttp.SendErrorResponse(c, http.StatusBadRequest, "validation error")
		}

		employeeId, err := strconv.ParseInt(id.EmployeeId, 10, 64)
		if err != nil {
			return
		}

		err = employeeService.UpdateEmployee(c, employeeId, req)
		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			thttp.SendErrorResponse(c, http.StatusNotFound, "employee not found")
			return
		default:
			logger.Error(err.Error())
			thttp.SendErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		thttp.SendOkResponse(c, nil)
	}
}
