package models

import (
	"time"
)

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(255)"`
	Items        []Item
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
