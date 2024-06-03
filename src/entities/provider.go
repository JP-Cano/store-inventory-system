package entities

import "github.com/google/uuid"

type Provider struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `json:"name"`
	ContactInfo string    `json:"contact_info"`
	Products    []Product
}
