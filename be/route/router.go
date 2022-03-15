package route

import (
	"be/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/students", api.GetStudents)
	e.POST("/students", api.PostStudent)
	e.PATCH("/students/:student_id", api.UpdateStudent)
	e.DELETE("/students/:student_id", api.DeleteStudent)

	e.GET("/courses", api.GetCourses)

	e.GET("/extras", api.GetExtras)

	return e
}
