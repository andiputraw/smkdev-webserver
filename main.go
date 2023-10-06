package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

var students []Student

func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

func getStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

func createStudent(c echo.Context) error {
	student := new(Student)

	if err := c.Bind(student); err != nil {
		return err
	}

	student.ID = len(students) + 1
	students = append(students, *student)

	return c.JSON(http.StatusCreated, student)
}

func updateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID != id {
			continue
		}
		updatedStudent := new(Student)
		if err := c.Bind(updatedStudent); err != nil {
			return err
		}
		students[i].Name = updatedStudent.Name
		students[i].Age = updatedStudent.Age
		students[i].Gender = updatedStudent.Gender
		students[i].ID = len(students) + 1

		return c.JSON(http.StatusOK, updatedStudent)
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

func deleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")

}

func main() {
	e := echo.New()
	e.GET("/students", getStudents)
	e.GET("/students/:id", getStudent)
	e.POST("/students", createStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)
	e.Logger.Fatal(e.Start(":1323"))
}
