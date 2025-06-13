package models

import (
	"errors"
	"time"

	"github.com/alfredamos/initializers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID        string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	ItemName        string         `gorm:"type:varchar(255)" json:"itemName"`
	Image        string         `gorm:"type:varchar(255)" json:"image"`
	Quantity     float64        `json:"quantity"`
	Price        float64        `json:"price"`
	OrderID  string `gorm:"foreignKey:OrderID;type:varchar(255)" json:"orderId"`
	Order Order `json:"order"`
	MenuItemID  string `gorm:"foreignKey:MenuItemID;type:varchar(255)" json:"menuItemId" binding:"required"`
}

// This functions are called before creating any Post
func (t *OrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

func (orderDetail *OrderDetail) CreateOrderDetail() error{
	//----> Insert the cart-item into the database.
	err := initializers.DB.Create(&orderDetail).Error 

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not created")
	}

	//----> Send back response
	return nil
}

func (*OrderDetail) DeleteOrderDetailById(id string) error{
	//----> Retrieve the cart-item with the given id.
	_, err := orderDetailGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not found")
	}

	//----> Delete the cart-item with the given id from database.
	err = initializers.DB.Unscoped().Delete(&OrderDetail{}, "id = ?", id).Error
	
	//----> Check for error.
	if err != nil {
		return errors.New("cart-item cannot be deleted")
	}

	//----> Send back the response.
	return nil
}

func (orderDetail *OrderDetail) EditOrderDetailId(id string) error{
	//----> Retrieve the cart-item with the given id.
	_, err := orderDetailGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not found")
	}

	//----> Update the cart-item in the database.
	err = initializers.DB.Model(&orderDetail).Updates(&orderDetail).Error

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item cannot be updated")
	}

	//----> Send back the response.
	return nil
}

func (*OrderDetail) GetAllOrderDetails() ([]OrderDetail, error){
	orderDetails := []OrderDetail{} //----> Declaration.

	//----> Retrieve the cart-items from the database.
	err := initializers.DB.Find(&orderDetails).Error

	//----> Check for error.
	if err != nil {
		return []OrderDetail{}, errors.New("cart-items are not found")
	}

	//----> send back the response.
	return orderDetails, nil
}

func (*OrderDetail) GetOrderDetailById(id string) (OrderDetail, error){
	//----> Retrieve the cart-item with the given id.
	orderDetail, err := orderDetailGetById(id)

	//----> Check for error.
	if err != nil {
		return OrderDetail{}, errors.New("cart-item is not found")
	}

	//----> send back the response.
	return orderDetail, nil

}
