package models

import "time"

// Item represents the model for an item
type Item struct {
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(255)"`
	Description string
	Quantity    uint
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
