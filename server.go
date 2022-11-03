package main

import (
	db "myapp/config"
	"myapp/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.InitialMigration()
	e := echo.New()
	// DB:=db.GetDBInstance()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/students", handler.GetAllStudent)
	e.GET("/students/:id", handler.GetStudent)
	e.POST("/students", handler.CreateStudent)
	e.PUT("/students/:id", handler.UpdateStudent)
	e.DELETE("/students/:id", handler.DeleteStudent)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
