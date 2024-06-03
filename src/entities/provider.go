package entities

import "github.com/google/uuid"

type Provider struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `json:"name"`
	ContactInfo string    `json:"contact_info"`
	Products    []Product `json:"-"`
}

type ProviderDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ContactInfo string    `json:"contact_info"`
}
