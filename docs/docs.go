// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/employee/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Создание работника",
                "parameters": [
                    {
                        "description": "Данные пользователя",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/test-task-sw_service_models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/employee/delete/{employeeId}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Удаление работника",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор работника",
                        "name": "employeeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseOk"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/employee/list/company/{companyId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Получение работников по id компании",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор компании",
                        "name": "companyId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseWithDetails-array_test-task-sw_service_models_Employee"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/employee/list/department/{depName}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Получение работников по отделу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название отдела",
                        "name": "depName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseWithDetails-array_test-task-sw_service_models_Employee"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/employee/update/{employeeId}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Изменение данных работника",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор работника",
                        "name": "employeeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления работника",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/test-task-sw_service_models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseWithDetails-test-task-sw_service_models_Employee"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    },
                    "409": {
                        "description": "Already exists",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/ping": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Служебные"
                ],
                "summary": "Пинг сервиса",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thttp.ResponseWithDetails-string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "test-task-sw_service_models.Department": {
            "type": "object",
            "properties": {
                "department_name": {
                    "type": "string"
                },
                "department_phone": {
                    "type": "string"
                }
            }
        },
        "test-task-sw_service_models.Employee": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "integer"
                },
                "department": {
                    "$ref": "#/definitions/test-task-sw_service_models.Department"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passport": {
                    "$ref": "#/definitions/test-task-sw_service_models.Passport"
                },
                "phone": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "test-task-sw_service_models.Passport": {
            "type": "object",
            "properties": {
                "passport_number": {
                    "type": "string"
                },
                "passport_type": {
                    "type": "string"
                }
            }
        },
        "thttp.ResponseError": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error_code": {
                    "type": "integer"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {}
            }
        },
        "thttp.ResponseOk": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error_code": {
                    "type": "integer"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {}
            }
        },
        "thttp.ResponseWithDetails-array_test-task-sw_service_models_Employee": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error_code": {
                    "type": "integer"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/test-task-sw_service_models.Employee"
                    }
                }
            }
        },
        "thttp.ResponseWithDetails-string": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error_code": {
                    "type": "integer"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "type": "string"
                }
            }
        },
        "thttp.ResponseWithDetails-test-task-sw_service_models_Employee": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error_code": {
                    "type": "integer"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/test-task-sw_service_models.Employee"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Test task Swartway",
	Description:      "API Server for application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
