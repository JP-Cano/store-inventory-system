package entities

import (
	"github.com/google/uuid"
	"time"
)

type Sale struct {
	Id           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	ProductId    uuid.UUID `gorm:"type:uuid;" json:"product_id"`
	CustomerName string    `json:"customer_name"`
	QuantitySold int       `json:"quantity_sold"`
	SaleDate     time.Time `json:"sale_date"`
	Product      *Product  `json:"product"`
}

type ProductSales struct {
	ProductId uuid.UUID `json:"product_id"`
	TotalSold int       `json:"total_sold"`
}

type SaleDto struct {
	Id           uuid.UUID    `json:"id"`
	CustomerName string       `json:"customer_name"`
	QuantitySold int          `json:"quantity_sold"`
	SaleDate     time.Time    `json:"sale_date"`
	Product      *ProductDto  `json:"product"`
	Provider     *ProviderDto `json:"provider"`
}
