package models

import (
	"time"
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID        string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	ItemName        string         `gorm:"type:varchar(255)" json:"itemName"`
	Quantity     float64        `json:"quantity"`
	Price        float64        `json:"price"`
	OrderID  string `gorm:"foreignKey:OrderID;type:varchar(255)" json:"orderId"`
	Order Order `json:"order"`
	MenuItemID  string `gorm:"foreignKey:MenuItemID;type:varchar(255)" json:"menuItemId" binding:"required"`
}