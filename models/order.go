package models

import (
	"errors"
	"time"
	"github.com/alfredamos/initializers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Confirmed Status = "Confirmed"
	Pickup Status = "Pickup"
	Cancelled Status = "Cancelled"
	Completed Status = "Completed"
)

type OrderPayload struct {
	UserId string `json:"userId"`
	PaymentId string `json:"paymentId"`
	OrderDetails []OrderDetail 
}


type Order struct {
	ID                string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	PaymentId 				string `json:"paymentId"`
	OrderDate         time.Time      `json:"orderDate"`
	Status            Status         `json:"status" binding:"required"`
	TotalQuantity     float64        `json:"totalQuantity"`
	TotalPrice        float64        `json:"totalPrice"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID" json:"orderDetails"`
}

// This functions are called before creating any Post
func (t *Order) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

func (order *Order) DeleteOrderById(id string) error{
	//----> Check to see if the order to be deleted is available in the database.
	err := initializers.DB.Model(&Order{}).Preload("CartItems").First(&order, "id = ?", id).Error
	
	//----> Check for error.
	if err != nil {
		return errors.New("order does not exist")
	}

	//----> Get the cart-items from order.
	carts := order.OrderDetails 

	//----> Delete all the cart-items attached to order with the given id.
	 err = deleteManyOrderDetails(carts)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-items cannot be deleted")
	} 
 
	//----> Delete the order with given id.
	err = initializers.DB.Unscoped().Delete(&Order{}, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return errors.New("order cannot be deleted")
	}

	return nil
}

func (*Order) DeleteOrderByUserId(userId string) error{
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Preload("CartItems").Find(&orders, Order{UserID: userId}).Error
	
	//----> Check for error.
	if err != nil {
		return errors.New("orders are not available in the database")
	}
	
	//----> Delete all orders and associated cart-items connected to this user-id.
	err = deleteManyOrders(orders)

	//----> Check for error.
	if err != nil {
		return errors.New("orders cannot be deleted")
	}

	return nil
}

func (*Order) DeleteAllOrders() error{
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Preload("CartItems").Find(&orders).Error
	
	//----> Check for error.
	if err != nil {
		return errors.New("orders are not available in the database")
	}
	
	//----> Delete all orders and associated cart-items connected to this user-id.
	err = deleteManyOrders(orders)

	//----> Check for error.
	if err != nil {
		return errors.New("orders cannot be deleted")
	}

	return nil
}

func (*Order) GetAllOrders() ([]Order, error){
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Model(&Order{}).Preload("User").Preload("CartItems").Find(&orders).Error

	//----> Check for error.
	if err != nil {
		return []Order{}, errors.New("orders are not available in the database")
	}

	//----> Send back response.
	return orders, nil
}

func (*Order) GetAllOrdersByUserId(userId string) ([]Order, error){
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Preload("User").Preload("CartItems").Find(&orders, Order{UserID: userId}).Error
	
	//----> Check for error.
	if err != nil {
		return []Order{}, errors.New("orders are not available in the database")
	}

	//----> Send back response.
	return orders, nil
}

func (order *Order) GetOrderById(id string) (Order, error){
	//----> retrieve the order with the given id from database.
	err := initializers.DB.Model(&Order{}).Preload("User").Preload("CartItems").First(&order, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return Order{}, errors.New("order is not available in the database ")
	}

	//----> Send back response.
	return *order, nil
}


func (order *OrderPayload) CheckOutOrder() error{
	//----> Get the carts slice.
	carts := order.OrderDetails //----> Cart-items.
	userId := order.UserId //----> User-id
	paymentId := order.PaymentId //----> Payment-id

	//----> Make order struct.
	orderPayload := makeOrder(userId, carts, paymentId)

	//----> Insert order in the database.
	err := initializers.DB.Create(&orderPayload).Error

	//----> Check for error.
	if err != nil{
		return errors.New("order creation fails")
	}

	//----> Get the orderPayload-id
  orderPayloadId := orderPayload.ID

	//----> Make cart-items from cart-item struct.
	cartItems := makeOrderDetails(carts, orderPayloadId)

	//----> Insert all the cart-items with the given order-id in the database.
	err = initializers.DB.CreateInBatches(&cartItems, len(cartItems)).Error

	//----> Check for error.
	if err != nil{
		return errors.New("cartItems creation fails")
	}

	return nil
}

