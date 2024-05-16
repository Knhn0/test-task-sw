package api

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net"
	"net/http"
	"test-task-sw/api/handler"
	_ "test-task-sw/docs"
	"test-task-sw/lib/tctx"
	"test-task-sw/service"
	"time"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	Server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(
	port int,
	logger *zap.SugaredLogger,
	contextProvider tctx.DefaultContextProviderFunc,
	employeeService *service.EmployeeService,
	passportService *service.PassportService,
	departmentService *service.DepartmentService,
) *Server {
	r := gin.New()

	//cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiGroup := r.Group("/api")
	{
		serviceGroup := apiGroup.Group("/")
		{
			serviceGroup.GET("/ping", handler.Ping())
		}

		employeeGroup := apiGroup.Group("/employee")
		{
			employeeGroup.POST("create", handler.CreateEmployee(logger, employeeService))
			employeeGroup.DELETE("delete/:employeeId", handler.DeleteEmployee(logger))
			employeeGroup.GET("list/:companyId", handler.ListEmployeesByCompanyId(logger))
			employeeGroup.GET("list/department/:depName", handler.ListEmployeesByDepartment(logger))
			employeeGroup.PUT("update/:employeeId", handler.UpdateEmployee(logger))
		}
	}
	return &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: r,
			BaseContext: func(listener net.Listener) context.Context {
				return contextProvider()
			},
		},
		shutdownTimeout: _defaultShutdownTimeout,
	}
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.Server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	return s.Server.Shutdown(shutdownCtx)
}
