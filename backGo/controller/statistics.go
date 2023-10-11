package controller

import (
	"fmt"
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
	fmt.Println("token: ", token)
	random_token := tools.GetTokenForStatistics(token, jsonData.Numbers, jsonData.Starttime, jsonData.Endtime)
	fmt.Println("random_token: ", random_token)
	stringCSV := tools.GetCSV(token, random_token, jsonData.Numbers, jsonData.Starttime, jsonData.Endtime)

	response := tools.SendCVStoProcess(stringCSV)

	context.JSON(http.StatusOK, gin.H{
		"inbounds":           response.Inbounds,
		"outbounds":          response.Outbounds,
		"inbounds_answered":  response.Inbounds_answered,
		"outbounds_answered": response.Outbounds_answered,
	})

}
