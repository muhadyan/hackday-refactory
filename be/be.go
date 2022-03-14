package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	connect()

	r.GET("/students", getStudents)
	r.POST("/students", postStudent)
	r.PATCH("/students/:student_id", updateStudent)
	r.DELETE("/students/:student_id", deleteStudent)

	r.GET("/courses", getCourses)

	r.GET("/extras", getExtras)

	r.Run()
}

type Students struct {
	StudentID int64  `json:"student_id" gorm:"not null;primaryKey;autoIncrement"`
	FullName  string `json:"full_name"`
	ExtraName string `json:"extra_name" gorm:"column:extra_name"`
}
// for post
type students struct {
	StudentID int64  `json:"student_id"`
	FullName  string `json:"full_name"`
	ExtraID   int64  `json:"extra_id"`
}
type Courses struct {
	CourseID   int64  `json:"course_id"`
	CourseName string `json:"course_name"`
}
type Extras struct {
	ExtraID   int64  `json:"extra_id"`
	ExtraName string `json:"extra_name"`
}

var DB *gorm.DB

func connect() {
	dsn := "root:test@tcp(127.0.0.1:3306)/hackday"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}

//// STUDENTS
// Get Data Students
func getStudents(c *gin.Context) {
	var students []Students
	DB.Table(`students`).
		Select(`students.student_id, students.full_name, extras.extra_name`).
		Joins("inner join extras on extras.extra_id = students.extra_id").
		Scan(&students)

	c.JSON(http.StatusOK, students)
}

// Post Data Students
func postStudent(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	// Validate Input
	var input students
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Validate Booking Id & Create Booking
	var check Students
	DB.Where("id = ?", input.StudentID).First(&check)

	student := students{StudentID: input.StudentID, FullName: input.FullName, ExtraID: input.ExtraID}
	DB.Create(&student)
	c.JSON(http.StatusOK, "Success")
}

// Update Data Students
func updateStudent(c *gin.Context) {
	// Get model if exist
	var student Students
	param := c.Param("student_id")

	if err := DB.Where("student_id = ?", param).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, "Record not found!")
	}

	// validate input
	var input Students
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	DB.Model(&student).Updates(input)

	c.JSON(http.StatusOK, student)
}

// Delete Data Students
func deleteStudent(c *gin.Context) {
	var student Students
	param := c.Param("student_id")

	if err := DB.Where("student_id = ?", param).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, "Record not found")
	}

	DB.Delete(&student)

	c.JSON(http.StatusOK, "Data has been deleted")
}

//// COURSES
// Get Data Courses
func getCourses(c *gin.Context) {
	var courses []Courses
	DB.Find(&courses)

	c.JSON(http.StatusOK, courses)
}

//// EXTRAS
// Get Data Extras
func getExtras(c *gin.Context) {
	var extras []Extras
	DB.Find(&extras)

	c.JSON(http.StatusOK, extras)
}
