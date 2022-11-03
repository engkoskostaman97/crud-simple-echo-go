package handler

import (
	"encoding/json"
	"fmt"
	db "myapp/config"
	"myapp/model"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func CreateStudent(c echo.Context) error {
	var student model.Student
	var DB *gorm.DB = db.GetDBInstance()
	err := json.NewDecoder(c.Request().Body).Decode(&student)
	if err != nil {
		log.Error("empty json body")
		return nil
	}
	DB.Create(&student)
	return c.JSON(http.StatusOK, student)
}

func GetStudent(c echo.Context) error {
	var student model.Student
	var DB *gorm.DB = db.GetDBInstance()
	id, _ := strconv.Atoi(c.Param("id"))
	DB.First(&student, id)
	return c.JSON(http.StatusOK, student)

}

func UpdateStudent(c echo.Context) error {

	var student model.Student
	var DB *gorm.DB = db.GetDBInstance()
	id, _ := strconv.Atoi(c.Param("id"))
	DB.First(&student, id)
	err := json.NewDecoder(c.Request().Body).Decode(&student)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	DB.Save(&student)
	return c.JSON(http.StatusOK, student)
}

func GetAllStudent(c echo.Context) error {
	var DB *gorm.DB = db.GetDBInstance()
	var students []model.Student
	DB.Find(&students)
	return c.JSON(http.StatusOK, students)
}

type SuccessRessponse struct {
	Success bool `json:"success"`
}

func DeleteStudent(c echo.Context) error {
	var DB *gorm.DB = db.GetDBInstance()
	id, _ := strconv.Atoi(c.Param("id"))
	var student model.Student
	DB.Delete(&student, id)

	r := &SuccessRessponse{
		Success: true,
	}
	return c.JSON(http.StatusOK, r)
}
