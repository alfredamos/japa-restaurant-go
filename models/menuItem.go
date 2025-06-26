package models

import (
	"errors"
	"time"
	"github.com/alfredamos/initializers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuItem struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	ItemName string `gorm:"type:varchar(255)" json:"itemName"`
	Category string `gorm:"type:varchar(255)" json:"category"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	SpecialTag string `gorm:"type:varchar(255)" json:"specialTag"`
	Image string `gorm:"type:varchar(255)" json:"image"`
	Price float64 `json:"price"`
	UserID string `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}

// This functions are called before creating any Post
func (t *MenuItem) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

func (menuItem *MenuItem) CreateMenuItem()(MenuItem, error){
	//----> Insert the menu-item in the database.
	err := initializers.DB.Create(&menuItem).Error

	//----> Check for insertion error.
	if err != nil {
		return MenuItem{}, errors.New("MenuItem cannot be inserted in the database")
	}

	//----> Send back the response.
	return *menuItem, nil
}

func (*MenuItem) DeleteMenuItemById(id string)error {
	//----> Retrieve menu-item with the given id from database.
	_, err := menuItemGetById(id)

	//----> Check for error in retrieval.
	if err != nil {
		return errors.New("menuItem cannot be retrieved from database")
	}

	//----> Delete the menu-item with the given id from database.
	err = initializers.DB.Unscoped().Delete(&MenuItem{}, "id = ?", id).Error

	//----> Check for deletion error.
	if err != nil {
		return errors.New("menu-item cannot be deleted")
	}

	//----> Send back the response.
	return nil
}

func (menuItem *MenuItem) EditMenuItemById(id string)error{
	//----> Retrieve menu-item with the given id from database.
	_, err := menuItemGetById(id)

	//----> Check for error in retrieval.
	if err != nil {
		return errors.New("menuItem cannot be retrieved from database")
	}

	//----> Update the menu-item in the database.
	err = initializers.DB.Model(&menuItem).Updates(&menuItem).Error

	//----> Check for update error.
	if err != nil {
		return errors.New("menuItem cannot be updated")
	}

	//-----> Send back the response
	return nil
}

func (*MenuItem) GetAllMenuItems() ([]MenuItem, error){
	//----> Declare the variable.
	menuItems := []MenuItem{}

	//----> Retrieve the menu-items from the database.
	err := initializers.DB.Find(&menuItems).Error

	//----> Check for retrieval error.
	if err != nil {
		return []MenuItem{}, errors.New("MenuItems cannot be retrieved")
	}

	//----> Send back the response.
	return menuItems, nil
}

func (*MenuItem) GetMenuItemById(id string)(MenuItem, error){
	//----> Retrieve menu-item with the given id from database.
	menuItem, err := menuItemGetById(id)

	//----> Check for error in retrieval.
	if err != nil {
		return MenuItem{}, errors.New("menuItem cannot be retrieved from database")
	}

	//-----> Send back the response.
	return menuItem, nil
}