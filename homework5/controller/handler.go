package controller

import (
	"homework5/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		EpTracker := model.MockData1
		context.JSON(200, EpTracker)
	}
}

func GetByIdExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		EpTrackerId := context.Param("id")
		IntEpTrackerId, err := strconv.Atoi(EpTrackerId)
		if err != nil {
			println(err)
		}
		context.JSON(200, model.MockData1[IntEpTrackerId-1])
	}
}

func PostExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data model.ExpenseTracker
		err := context.ShouldBindJSON(&data)
		if err != nil {
			println(err)
		}
		model.MockData1 = append(model.MockData1, data)
	}
}

func PutExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramID := context.Param("id")
		IntParamId, err1 := strconv.Atoi(paramID)
		if err1 != nil {
			println(err1)
		}
		var data model.ExpenseTracker
		err2 := context.ShouldBindJSON(&data)
		if err2 != nil {
			println(err2)
		}
		model.MockData1[IntParamId-1] = data
	}
}

func DeleteExpenseTrackerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		param := context.Param("id")
		paramID, err:= strconv.Atoi(param)
		if err != nil {
			println(err)
		}
		for i := 0; i <= len(model.MockData1) - 1; i++ {
			if model.MockData1[i].ID == paramID {     
				model.MockData1 = append(model.MockData1[:i], model.MockData1[i+1:]...)
					   context.JSON(200, "delete success")
					   
					   return
					 }
		}
		context.JSON(http.StatusNotFound, "Data not found")
	}
}
