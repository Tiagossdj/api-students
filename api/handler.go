package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/Tiagossdj/api-students/db"
	"github.com/labstack/echo/v4"
)

func(api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil{
		return c.String(http.StatusNotFound, "Failed to get students!")
	}
  return c.JSON(http.StatusOK, students)

}

func(api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
			return err
	}

	if err := api.DB.AddStudent(student); err != nil {
  	return c.String(http.StatusInternalServerError, "Error to Create student")
	} 

	return c.String(http.StatusOK, "Create student")
	}

func(api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student id")
	}

student,err := api.DB.GetStudent(id)
if errors.Is(err, gorm.ErrRecordNotFound){
return c.String(http.StatusNotFound, "Student not found")
}

if err != nil {
	return c.String(http.StatusInternalServerError, "Failed to get Student")
	}
 return c.JSON(http.StatusOK,student)
}

func(api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update %s student", id)
  return c.String(http.StatusOK, updateStud)
}

func(api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete %s student", id)
  return c.String(http.StatusOK, deleteStud)
}