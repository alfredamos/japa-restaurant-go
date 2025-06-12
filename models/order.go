package models

import (
	"time"
	"gorm.io/gorm"
)

type Status string

const (
	Confirmed Status = "Confirmed"
	Pickup Status = "Pickup"
	Cancelled Status = "Cancelled"
)

type Order struct {
	ID                string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	PickupName        string         `gorm:"type:varchar(255)" json:"pickupName"`
	PickupEmail       string         `gorm:"unique;type:varchar(255)" json:"pickupEmail"`
	PickupPhoneNumber string         `gorm:"type:varchar(255)" json:"pickupPhoneNumber"`
	OrderDate         time.Time      `json:"orderDate"`
	Status            Status         `json:"status" binding:"required"`
	TotalQuantity     float64        `json:"totalQuantity"`
	TotalPrice        float64        `json:"totalPrice"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID" json:"orderDetails"`
}