package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"store-inventory-management/src/entities"
)

type ProductRepository interface {
	GetAll() ([]entities.Product, error)
	Search(value string) ([]entities.Product, error)
	UpdateQuantity(id uuid.UUID, quantity int) error
	GetById(id uuid.UUID) (*entities.Product, error)
	ValidateProductForSale(id uuid.UUID, quantitySold int) (*entities.Product, error)
}

type ProductStorage struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) *ProductStorage {
	return &ProductStorage{
		db: db,
	}
}

func (p *ProductStorage) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	if err := p.db.Preload("Provider").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductStorage) Search(value string) ([]entities.Product, error) {
	var products []entities.Product
	searchPattern := "%" + value + "%"
	query := p.db.Table("products").
		Select("products.*, providers.name as provider_name").
		Joins("JOIN providers ON providers.id = products.provider_id").
		Where("products.name ILIKE ? OR products.category ILIKE ? OR providers.name ILIKE ?", searchPattern, searchPattern, searchPattern).
		Preload("Provider")

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductStorage) UpdateQuantity(id uuid.UUID, quantity int) error {
	product, err := p.GetById(id)

	if err != nil {
		return err
	}

	if err = p.db.Model(&product).Update("quantity_available", gorm.Expr("quantity_available - ?", quantity)).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductStorage) GetById(id uuid.UUID) (*entities.Product, error) {
	var product entities.Product
	if err := p.db.Preload("Provider").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductStorage) ValidateProductForSale(id uuid.UUID, quantitySold int) (*entities.Product, error) {
	product, err := p.GetById(id)
	if err != nil {
		return nil, err
	}

	if product.QuantityAvailable < 0 {
		return nil, errors.New("product quantity is not available")
	}

	if product.QuantityAvailable < quantitySold {
		return nil, errors.New("insufficient product quantity")
	}

	return product, nil
}
