package models

import (
	"time"
	"gorm.io/gorm"
)

type MenuItem struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Category string `gorm:"type:varchar(255)" json:"category"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	SpecialTag string `gorm:"type:varchar(255)" json:"specialTag"`
	Image string `gorm:"type:varchar(255)" json:"image"`
	Price float64 `json:"price"`
	UserID string `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
	User User `json:"user"` 
}