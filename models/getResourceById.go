package models

import (
	"errors"

	"github.com/alfredamos/initializers"
)

func orderDetailGetById(id string) (OrderDetail, error) {
	orderDetail := OrderDetail{} //----> Declaration.

	//----> Retrieve the cart-item with given id from database.
	err := initializers.DB.First(&orderDetail, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return OrderDetail{}, errors.New("the cart-item with the given id is not found")
	}

	//----> Send back the response.
	return orderDetail, nil
}


func menuItemGetById(id string) (MenuItem, error) {
	menuItem := MenuItem{} //----> MenuItem variable.

	//----> Retrieve the menuItem with the given id from the database.
	err := initializers.DB.First(&menuItem, "id = ?", id).Error

	//----> Check for non existent menuItem.
	if err != nil {
		return MenuItem{}, errors.New("the menuItem with the given id is not found")
	}

	//----> Send back the response.
	return menuItem, nil
}

func userGetById(id string) (User, error) {
	user := User{} //----> User variable.
	
	//----> Retrieve the user with the given id from the database.
	err := initializers.DB.Omit("Password").First(&user, "id = ?", id).Error

	//----> Check for non existent user.
	if err != nil {
		return User{}, errors.New("there is no user with the given id to retrieve from database")
	}

	//----> Send back the response.
	return user, nil
}

func getAllOrderDetailsIds(carts []OrderDetail)[]OrderDetail{
	orderDetails := make([]OrderDetail, 0) //----> Slice of cart-ids
	
	//----> Get all the cart-items ids.
	for _, cart := range carts {
		//----> Compose the id from the cart-item struct.
		orderDetail := OrderDetail{ ID: cart.ID}

		//----> Append all the ids together to have a slice of cart-item ids.
		orderDetails = append(orderDetails, orderDetail)

	}

	//----> Send back the response
	return orderDetails
}
 