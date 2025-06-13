package models

import (
	"errors"
	"time"
	"github.com/alfredamos/initializers"
)

func CalTotalPriceAndQuantity(carts []OrderDetail) (float64, float64) {
	//----> Initialize totalQuantity and totalPrice.
	totalQuantity := 0.0 //----> Total quantity
	totalPrice := 0.0 //----> Total price.

	//----> Calculate the totalQuantity and totalPrice.
	for _, value := range carts {
		//----> Total quantity.
		totalQuantity += value.Quantity

		//----> Total price.
		totalPrice += value.Quantity * value.Price
	}

	return totalQuantity, totalPrice
}

func makeOrder(userId string, carts []OrderDetail, paymentId string) Order{
	//----> Get the total quantity and total price.
	totalQuantity, totalPrice := CalTotalPriceAndQuantity(carts)

	//----> Make order.
	order := Order{
		UserID:        userId,
		PaymentId:     getPaymentId(paymentId),
		OrderDate:     time.Now(),
		TotalQuantity: totalQuantity,
		TotalPrice:    totalPrice,
		Status:        "Pending",
	}

	return order
}

func getPaymentId(paymentId string) string{
	result := string("") //----> Initialize result.

	//----> Set payment id as appropriate.
	if len(paymentId) == 0 {
			result = "wehdkjrifbsvss" //----> Payment-id is not given.
	} else {
			result = paymentId //----> Payment-id is given.
	}

	return result
}

func makeOrderDetails(carts []OrderDetail, orderId string) []OrderDetail {
	newCarts := []OrderDetail{} //----> Cart variable.

	//----> Make the cart-items by composing cart-item struct.
	for _, value := range carts {
		newCart := OrderDetail{
			ItemName:     value.ItemName,
			Price:    value.Price,
			Quantity: value.Quantity,
			Image:    value.Image,
			OrderID:  orderId,
			MenuItemID:  value.MenuItemID,
		}

		//----> Append newCart to newCarts.
		newCarts = append(newCarts, newCart)
	}

	return newCarts
}

func deleteManyOrderDetails(carts []OrderDetail) error{
	//----> Get all the ids of the cart-items to be deleted.
	orderDetails := getAllOrderDetailsIds(carts)
	
	//----> Delete all cart-items.
	err := initializers.DB.Unscoped().Delete(&orderDetails).Error

	//----> Check for error.
	if err != nil {
		return errors.New("cart-items cannot be deleted")
	}

	return nil
}

func deleteManyOrders(orders []Order) error{
	allOrders := make([]Order,0) //----> orders - ids.
	
	//----> Get ids of orders to be deleted.
	for _, order := range orders{
		oneOrder := Order{ID: order.ID}//----> Order-id.
		allOrders  = append(allOrders , oneOrder) //----> orders-ids.

		//----> Delete all cart-items associated with this order.
		err := deleteManyOrderDetails(order.OrderDetails)

		if err != nil{
			return errors.New("cart-items cannot be deleted")
		}
	}

	//----> Delete all orders.
	err := initializers.DB.Unscoped().Delete(&allOrders).Error

	//----> Check for error.
	if err != nil {
		return errors.New("orders cannot be deleted")
	}

	return nil
}

