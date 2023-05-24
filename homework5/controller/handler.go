package controller

import (
	"homework5/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		expenseTracker := model.MockData1
		context.JSON(200, expenseTracker)
	}
}

func GetByIdExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		trackerIdStr := context.Param("id")
		trackerId, err := strconv.Atoi(trackerIdStr)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, "Incorrect ID")
		}
		context.JSON(200, model.MockData1[trackerId-1])
	}
}

func PostExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data model.ExpenseTracker
		err := context.ShouldBindJSON(&data)
		if err != nil {
			context.JSON(404, "Data not found")
		}
		model.MockData1 = append(model.MockData1, data)
	}
}

func PutExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		trackerIdStr := context.Param("id")
		trackerId, err1 := strconv.Atoi(trackerIdStr)
		if err1 != nil {
			context.JSON(http.StatusNotAcceptable, "Incorrect ID")
		}
		var data model.ExpenseTracker
		err2 := context.ShouldBindJSON(&data)
		if err2 != nil {
			context.JSON(404, "Data not found")
		}
		model.MockData1[trackerId-1] = data
	}
}

func DeleteExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		trackerIdStr := context.Param("id")
		trackerId, err := strconv.Atoi(trackerIdStr)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, "Incorrect ID")
		}
		for i := 0; i <= len(model.MockData1)-1; i++ {
			if model.MockData1[i].ID == trackerId {
				model.MockData1 = append(model.MockData1[:i], model.MockData1[i+1:]...)
				context.JSON(200, "delete success")

				return
			}
		}
		context.JSON(http.StatusNotFound, "Data not found")
	}
}
