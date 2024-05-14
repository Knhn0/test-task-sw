package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"test-task-sw/entity"
	"test-task-sw/lib/thttp"
	"test-task-sw/service"
)

type createUserRequest struct {
	Id              int    `json:"id" binding:"required" example:"1"`
	Name            string `json:"name" binding:"required" example:"Ivan"`
	Surname         string `json:"surname" binding:"required" example:"Ivanov"`
	Phone           string `json:"phone" binding:"required" example:"+7(987)6667788"`
	CompanyId       int    `json:"company_id" binding:"required" example:"1"`
	PassportName    string `json:"passport_name" binding:"required" example:"Russian passport"`
	PassportNumber  string `json:"passport_number" binding:"required" example:"1122 112233"`
	DepartmentName  string `json:"department_name" binding:"required" example:"First Department"`
	DepartmentPhone string `json:"department_phone" binding:"required" example:"+7(987)1112233"`
}

// CreateUser godoc
// @Summary     Созданеие пользователя
// @Tags		Users
// @Accept      json
// @Produce     json
// @Param       request body createUserRequest true "Данные пользователя"
// @Success     200
// @Security    ApiKeyAuth
// @Router      /api/users/create [post]
func CreateUser(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createUserRequest
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		var err error
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
		thttp.OkResponseWithResult(c, 1)
		return
	}
}

type deleteUserUri struct {
	UserId string `uri:"userId"`
}

// DeleteUser godoc
// @Summary     Удаление пользователя
// @Tags		Users
// @Accept      json
// @Produce     json
// @Param       userId path string true "Идентификатор пользователя"
// @Success     200 {object} thttp.ResponseOk
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Security    ApiKeyAuth
// @Router      /api/users/delete/{userId} [delete]
func DeleteUser(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req deleteUserUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		//err := profileService.Delete(c, uuid.MustParse(req.UserId))
		var err error
		switch {
		case err == nil:
		default:
			logger.Error(err.Error())
			thttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		thttp.OkResponse(c)
	}
}

type listUserByCompanyIdUri struct {
	CompanyId string `json:"companyId"`
}

type listUsersByCompanyIdResponse []listUsersByCompanyIdElement

type listUsersByCompanyIdElement struct {
	Id              int    `json:"id" binding:"required" example:"1"`
	Name            string `json:"name" binding:"required" example:"Ivan"`
	Surname         string `json:"surname" binding:"required" example:"Ivanov"`
	Phone           string `json:"phone" binding:"required" example:"+7(987)6667788"`
	CompanyId       int    `json:"company_id" binding:"required" example:"1"`
	PassportName    string `json:"passport_name" binding:"required" example:"Russian passport"`
	PassportNumber  string `json:"passport_number" binding:"required" example:"1122 112233"`
	DepartmentName  string `json:"department_name" binding:"required" example:"First Department"`
	DepartmentPhone string `json:"department_phone" binding:"required" example:"+7(987)1112233"`
}

func newListUsersByCompanyIdResponse(users []entity.User) listUsersByCompanyIdResponse {
	var response = make(listUsersByCompanyIdResponse, 0, len(users))
	for _, user := range users {
		response = append(response, listUsersByCompanyIdElement{
			Id:              user.Id,
			Name:            user.Name,
			Surname:         user.Surname,
			Phone:           user.Phone,
			CompanyId:       user.CompanyId,
			PassportName:    user.Passport.Name,
			PassportNumber:  user.Passport.Number,
			DepartmentName:  user.Department.Name,
			DepartmentPhone: user.Department.Phone,
		})
	}
	return response
}

// ListUsersByCompanyId godoc
// @Summary     Получение пользователей по id компании
// @Tags		Users
// @Accept      json
// @Produce     json
// @Param       companyId path string true "Идентификатор компании"
// @Success     200 {object} thttp.ResponseWithDetails[listUsersByCompanyIdResponse]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Security    ApiKeyAuth
// @Router      /api/users/list/{companyId} [get]
func ListUsersByCompanyId(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listUserByCompanyIdUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		var err error
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

		response := 1
		thttp.OkResponseWithResult(c, response)
	}
}

type listUsersByDepartmentUri struct {
	DepName string `uri:"depName"`
}

type listUsersByDepartmentResponse []listUsersByDepartmentElement

type listUsersByDepartmentElement struct {
	Id              int    `json:"id" binding:"required" example:"1"`
	Name            string `json:"name" binding:"required" example:"Ivan"`
	Surname         string `json:"surname" binding:"required" example:"Ivanov"`
	Phone           string `json:"phone" binding:"required" example:"+7(987)6667788"`
	CompanyId       int    `json:"company_id" binding:"required" example:"1"`
	PassportName    string `json:"passport_name" binding:"required" example:"Russian passport"`
	PassportNumber  string `json:"passport_number" binding:"required" example:"1122 112233"`
	DepartmentName  string `json:"department_name" binding:"required" example:"First Department"`
	DepartmentPhone string `json:"department_phone" binding:"required" example:"+7(987)1112233"`
}

func newListUsersByDepartmentResponse(users []entity.User) listUsersByDepartmentResponse {
	var response = make(listUsersByDepartmentResponse, 0, len(users))
	for _, user := range users {
		response = append(response, listUsersByDepartmentElement{
			Id:              user.Id,
			Name:            user.Name,
			Surname:         user.Surname,
			Phone:           user.Phone,
			CompanyId:       user.CompanyId,
			PassportName:    user.Passport.Name,
			PassportNumber:  user.Passport.Number,
			DepartmentName:  user.Department.Name,
			DepartmentPhone: user.Department.Phone,
		})
	}
	return response
}

// ListUsersByDepartment godoc
// @Summary     Получение пользователей по отделу
// @Tags		Users
// @Accept      json
// @Produce     json
// @Param       depName path string true "Название отдела"
// @Success     200 {object} thttp.ResponseWithDetails[listUsersByDepartmentResponse]
// @Failure     400 {object} thttp.ResponseError "Bad request"
// @Failure     500 {object} thttp.ResponseError "Internal server error"
// @Security    ApiKeyAuth
// @Router      /api/users/list/department/{depName} [get]
func ListUsersByDepartment(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listUsersByDepartmentUri
		if err := c.ShouldBindUri(&req); err != nil {
			thttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		var err error
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

		response := 1
		thttp.OkResponseWithResult(c, response)
	}
}
