package entities

import "github.com/google/uuid"

type Product struct {
	Id                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name              string    `json:"name"`
	Category          string    `json:"category"`
	QuantityAvailable int       `json:"quantity_available"`
	Price             float64   `json:"price"`
	ProviderId        uuid.UUID `gorm:"type:uuid;" json:"provider_id"`
	Provider          Provider
}
