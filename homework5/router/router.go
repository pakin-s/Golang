package router

import (
	"homework5/controller"

	"github.com/gin-gonic/gin"
)

func Init() {

	router := gin.Default()

	router.GET("/expense",
	controller.GetAllExpenseTrackerHandler())

	router.GET("/expense/:id",
	controller.GetByIdExpenseTrackerHandler())

	router.POST("/expense/Post",
	controller.PostExpenseTrackerHandler())

	router.PUT("/expense/Put/:id",
	controller.PutExpenseTrackerHandler())
	
	router.DELETE("/expense/Delete/:id",
	controller.DeleteExpenseTrackerHandler())

	router.Run(":8092")
}
