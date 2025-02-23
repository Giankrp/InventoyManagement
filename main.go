package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
	"github.com/Giankrp/inventoryManagement/routes"
)

type CustomValidator struct{
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error{
	return cv.validator.Struct(i)
}

func main() {
	db.DbConnection()
	db.Db.AutoMigrate(&models.Product{},&models.Brand{})
	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  e.Validator = &CustomValidator{validator: models.NewValidator()}


	routes.Routes(e)
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start server ", "error ", err)
	}
}
