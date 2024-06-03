package server

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Api struct {
	Port string
	DB   *gorm.DB
}

func New(port string, db *gorm.DB) *Api {
	return &Api{
		Port: port,
		DB:   db,
	}
}

func (api *Api) Serve() {
	e := echo.New()
	e.Logger.Fatal(e.Start(api.Port))
}
