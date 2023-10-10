package tools

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetToken() string {
	username := os.Getenv("USERNAME_CDR")
	password := os.Getenv("PASSWORD_CDR")
	port := os.Getenv("PORT_CDR")
	version := os.Getenv("VERSION_CDR")

	postBody, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
		"port":     port,
		"version":  version,
	})

	responseBody := bytes.NewReader(postBody)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post("https://172.25.31.100:8088/api/v1.10/login", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Response struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}

	var response Response

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("error:", err)
	}

	return response.Token
}

func updateToken(token string) string {

	params := url.Values{}
	params.Add("token", token)
	url := "https://172.25.31.100:8088/api/v1.10/heartbeat?" + params.Encode()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sp := string(body)

	return sp
}

func GetTokenForStatistics(token string, number string, starttime string, endtime string) string {

	params := url.Values{}
	params.Add("token", token)

	postBody, _ := json.Marshal(map[string]string{
		"number":    number,
		"starttime": starttime,
		"endtime":   endtime,
	})

	sendBody := bytes.NewReader(postBody)

	url := "https://172.25.31.100:8088/api/v2.0.0/cdr/get_random?" + params.Encode()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post(url, "application/json", sendBody)

	if err != nil {
		fmt.Println("error", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error", err)
	}

	type ResponseRandom struct {
		Status    string `json:"status"`
		Number    string `json:"number"`
		Starttime string `json:"starttime"`
		Endtime   string `json:"endtime"`
		Random    string `json:"random"`
	}

	var response ResponseRandom

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("error:", err)
	}

	return response.Random
}

func GetCSV(token string, random_token string, number string, starttime string, endtime string) {
	params := url.Values{}
	params.Add("number", number)
	params.Add("starttime", starttime)
	params.Add("endtime", endtime)
	params.Add("token", token)
	params.Add("random", random_token)

	url := "https://172.25.31.100:8088/api/v2.0.0/cdr/download?" + params.Encode()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", err)
	}

	stringBody := string(body)

	f, err := os.Create("statistic.csv")

	if err != nil {
		fmt.Println("error: ", err)
	}

	defer f.Close()

	_, err2 := f.WriteString(stringBody)

	if err2 != nil {
		fmt.Println("error: ", err2)
	}
}
