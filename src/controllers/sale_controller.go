package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"store-inventory-management/src/entities"
	"store-inventory-management/src/services"
)

type SaleController struct {
	saleRepository services.SaleRepository
}

func NewSaleController(salesRepository services.SaleRepository) *SaleController {
	return &SaleController{
		saleRepository: salesRepository,
	}
}

func (s *SaleController) RegisterRoutes(e *echo.Echo) {
	router := e.Group("/sales")
	router.GET("/search", s.findByDate)
	router.POST("", s.register)
	router.GET("/report", s.getProductSalesBetweenDates)
}

func (s *SaleController) register(e echo.Context) error {
	var newSale entities.Sale
	if err := e.Bind(&newSale); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	sale, err := s.saleRepository.Register(newSale)
	if err != nil {
		e.Logger().Error(err)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusCreated, map[string]interface{}{"data": sale})
}

func (s *SaleController) findByDate(e echo.Context) error {
	date := e.QueryParam("date")
	sales, err := s.saleRepository.FindByDate(date)
	if err != nil {
		e.Logger().Error(err)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search sale"})
	}

	return e.JSON(http.StatusOK, map[string][]entities.Sale{"data": sales})
}

func (s *SaleController) getProductSalesBetweenDates(e echo.Context) error {
	startDate := e.QueryParam("start-date")
	endDate := e.QueryParam("end-date")

	productSales, err := s.saleRepository.GetProductSalesBetweenDates(startDate, endDate)
	if err != nil {
		e.Logger().Error(err)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate report by dates"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"data": productSales})
}
