package routes

import (
	"github.com/alfredamos/controllers"
	"github.com/gin-gonic/gin"
)

func protectedRoutes(r gin.IRoutes){
	//----> Auth routes.
	r.GET("/auth/current-user", controllers.GetCurrentUserController)
	r.PATCH("/auth/change-password", controllers.ChangePasswordController)
	r.PATCH("/auth/edit-profile", controllers.EditProfileController)
	r.POST("/auth/logout",controllers.LogoutController)
	
	//----> Order-detail routes.
	r.GET("/order-details", controllers.GetAllOrderDetails)
	r.POST("/order-details", controllers.CreateOrderDetail)
	r.DELETE("/order-details/:id", controllers.DeleteOrderDetailById)
	r.GET("/order-details/:id", controllers.GetOrderDetailById)
	r.PATCH("/order-details/:id", controllers.EditOrderDetailById)

	//----> Order routes.
	r.PATCH("/orders/checkout", controllers.CheckOutOrder)
	
	//----> Menu-item-routes.
	r.GET("/menu-items/:id", controllers.GetMenuItemById)

	//----> Stripe payment-route
	r.POST("/stripe-payment/checkout", controllers.CreatePaymentController)

}