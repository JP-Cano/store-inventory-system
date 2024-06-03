package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"store-inventory-management/src/entities"
	"time"
)

type SaleRepository interface {
	Register(sale entities.Sale) (*entities.SaleDto, error)
	FindByDate(date string) ([]entities.Sale, error)
	GetProductSalesBetweenDates(startDate, endDate string) ([]entities.ProductSales, error)
}

type SaleStorage struct {
	db                *gorm.DB
	productRepository ProductRepository
}

func NewSaleStorage(db *gorm.DB, productRepository ProductRepository) *SaleStorage {
	return &SaleStorage{
		db:                db,
		productRepository: productRepository,
	}
}

func (s *SaleStorage) Register(sale entities.Sale) (*entities.SaleDto, error) {
	sale.Id = uuid.New()
	sale.SaleDate = time.Now()

	product, err := s.productRepository.ValidateProductForSale(sale.ProductId, sale.QuantitySold)

	if err != nil {
		return nil, err
	}

	if err = s.db.Create(&sale).Error; err != nil {
		return nil, err
	}

	if err = s.productRepository.UpdateQuantity(sale.ProductId, sale.QuantitySold); err != nil {
		return nil, err
	}

	return s.mapToSaleDto(sale, product), nil
}

func (s *SaleStorage) mapToSaleDto(sale entities.Sale, product *entities.Product) *entities.SaleDto {
	return &entities.SaleDto{
		Id:           sale.Id,
		CustomerName: sale.CustomerName,
		QuantitySold: sale.QuantitySold,
		SaleDate:     sale.SaleDate,
		Product: &entities.ProductDto{
			Id:       product.Id,
			Name:     product.Name,
			Category: product.Category,
			Price:    product.Price,
		},
		Provider: &entities.ProviderDto{
			Id:          product.Provider.Id,
			Name:        product.Provider.Name,
			ContactInfo: product.Provider.ContactInfo,
		},
	}
}

func (s *SaleStorage) FindByDate(date string) ([]entities.Sale, error) {
	var sales []entities.Sale

	query := s.db.Table("sales").
		Select("sales.id, sales.customer_name, sales.quantity_sold, sales.sale_date, products.id as product_id").
		Joins("JOIN products ON products.id = sales.product_id").
		Where("DATE(sales.sale_date) = ?", date)

	if err := query.Scan(&sales).Error; err != nil {
		return nil, err
	}

	for i := range sales {
		product, err := s.productRepository.GetById(sales[i].ProductId)
		if err != nil {
			return nil, err
		}
		sales[i].Product = product
	}

	return sales, nil
}

func (s *SaleStorage) GetProductSalesBetweenDates(startDate, endDate string) ([]entities.ProductSales, error) {
	var productSales []entities.ProductSales

	query := s.db.Table("sales").
		Select("product_id, SUM(quantity_sold) AS total_sold").
		Where("DATE(sale_date) BETWEEN ? AND ?", startDate, endDate).
		Group("product_id").
		Order("total_sold DESC")

	if err := query.Find(&productSales).Error; err != nil {
		return nil, err
	}

	return productSales, nil
}
