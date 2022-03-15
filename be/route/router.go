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
	e.POST("/courses", api.PostCourse)
	e.PATCH("/courses/:course_id", api.UpdateCourse)
	e.DELETE("/courses/:course_id", api.DeleteCourse)

	e.GET("/extras", api.GetExtras)
	e.POST("/extras", api.PostExtra)
	e.PATCH("/extras/:extra_id", api.UpdateExtra)
	e.DELETE("/extras/:extra_id", api.DeleteExtra)

	return e
}
