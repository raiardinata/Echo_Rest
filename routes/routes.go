package routes

import (
	"net/http"

	"echo_rest/controllers"
	"echo_rest/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Initialize Echo Framework!")
	})

	e.GET("/_health", controllers.HealthCheck)

	e.GET("/employee", controllers.FetchAllEmployee, middleware.IsAuthenticated)
	e.POST("/employee", controllers.PostEmployee, middleware.IsAuthenticated)
	e.PUT("/employee", controllers.PutEmployee, middleware.IsAuthenticated)
	e.DELETE("/employee", controllers.DelEmployee, middleware.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("/test-struct-validation", controllers.TestStructValidation)
	e.GET("/test-variable-validation", controllers.TestVariableValidation)

	return e
}
