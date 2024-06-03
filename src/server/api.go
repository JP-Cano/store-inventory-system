package server

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"store-inventory-management/src/controllers"
	"store-inventory-management/src/services"
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

func (a *Api) Serve() {
	e := echo.New()
	a.registerDependencies(e, a.DB)
	e.Logger.Fatal(e.Start(a.Port))
}

func (a *Api) registerDependencies(e *echo.Echo, db *gorm.DB) {
	// Product
	productRepository := services.NewProductStorage(db)
	productController := controllers.NewProductController(productRepository)
	productController.RegisterProductRoutes(e)

	// Sales
	saleRepository := services.NewSaleStorage(db, productRepository)
	saleController := controllers.NewSaleController(saleRepository)
	saleController.RegisterRoutes(e)
}
