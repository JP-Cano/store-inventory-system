package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"store-inventory-management/src/services"
)

type ProductController struct {
	productRepository services.ProductRepository
}

func NewProductController(productRepository services.ProductRepository) *ProductController {
	return &ProductController{
		productRepository: productRepository,
	}
}

func (p *ProductController) RegisterProductRoutes(e *echo.Echo) {
	router := e.Group("/products")
	router.GET("", p.getAll)
	router.GET("/search", p.search)
}

func (p *ProductController) getAll(e echo.Context) error {
	products, err := p.productRepository.GetAll()
	if err != nil {
		e.Logger().Error(err)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve all products"})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{"data": products})
}

func (p *ProductController) search(e echo.Context) error {
	value := e.QueryParam("value")
	products, err := p.productRepository.Search(value)
	if err != nil {
		e.Logger().Error(err)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search a product"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"data": products})
}
