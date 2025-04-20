package models

import (
	"github.com/google/uuid"
)

type Exercise struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
}
