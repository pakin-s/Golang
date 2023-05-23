package router

import (
	"homework5/controller"

	"github.com/gin-gonic/gin"
)

func Init() {

	router := gin.Default()

	router.GET("/expense", controller.GetAllExpenseTrackerHandler())
	router.GET("/expense/:id", controller.GetByIdExpenseTrackerHandler())
	router.POST("/expensePost", controller.PostExpenseTrackerHandler())
	router.PUT("/expensePut/:id", controller.PutExpenseTrackerHandler())
	router.DELETE("/expenseDelete/:id", controller.DeleteExpenseTrackerHandler())

	router.Run(":8092")
}
