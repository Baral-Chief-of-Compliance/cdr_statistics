package controller

import (
	"net/http"

	"github.com/cdr/tools"
	"github.com/gin-gonic/gin"
)

type JsonData struct {
	Numbers   string `json:"numbers"`
	Starttime string `json:"starttime"`
	Endtime   string `json:"endtime"`
}

func GetStatistics(context *gin.Context) {
	var jsonData JsonData

	context.BindJSON(&jsonData)

	token := tools.GetToken()
	random_token := tools.GetTokenForStatistics(token, jsonData.Numbers, jsonData.Starttime, jsonData.Endtime)
	tools.GetCSV(token, random_token, jsonData.Numbers, jsonData.Starttime, jsonData.Endtime)

	context.JSON(http.StatusUnauthorized, gin.H{
		"message": "da",
	})

}
