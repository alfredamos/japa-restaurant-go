package controllers

import (
	"net/http"
	"github.com/alfredamos/models"
	"github.com/gin-gonic/gin"
)

func CreateMenuItem(context *gin.Context){
	menuItem := models.MenuItem{} //----> Initialize the payload.
	
	//----> Get menu-payload.
	err := context.ShouldBind(&menuItem)

	//----> Check for bad request error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//----> Insert the menu-item into the database.
	newMenuItem, err := menuItem.CreateMenuItem()

	//----> Check for creation error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusCreated, newMenuItem)
}

func DeleteMenuItemById(context *gin.Context){
	menuItem := models.MenuItem{} //----> Initialize menu-item.
	
	//----> Get the menu-item id.
	id := context.Param("id")

	//----> Delete the menu-item in the database.
	err := menuItem.DeleteMenuItemById(id)

	//----> Check for deletion error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "MenuItem is deleted successfully!"})
}

func EditMenuItemById(context *gin.Context){
	menuItem := models.MenuItem{} //----> Initialize menu-item.
	
	//----> Get the menu-item to be edited from the payload.
	err := context.ShouldBind(&menuItem)

	//----> Check for bad request error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//----> Get the menu-item id.
	id := context.Param("id")

	//----> Update the menu-item in the database.
	err = menuItem.EditMenuItemById(id)

	//----> Check for update error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "MenuItem is updated successfully!"})
	
}

func GetAllMenuItems(context *gin.Context){
	menuItem := models.MenuItem{} //----> Initialize menu-item.

	//----> Get the menu-items from the database.
	menuItems, err := menuItem.GetAllMenuItems()

	//----> Check for retrieval error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, menuItems)
}

func GetMenuItemById(context *gin.Context){
	menuItem := models.MenuItem{} //----> Initialize menu-item.
	
	//----> Get the menu-item id.
	id := context.Param("id")

	//----> Get the menu-item with the given id from database.
	retrievedMenuItem, err := menuItem.GetMenuItemById(id)

	//----> Check for retrieval error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, retrievedMenuItem)
}