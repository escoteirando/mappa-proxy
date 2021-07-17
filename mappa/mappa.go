package mappa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const UrlMappa = "http://mappa.escoteiros.org.br"

var HttpClient = &http.Client{}

func GetRequest(c *gin.Context) {
	tudo := c.Param("request")
	url := UrlMappa + tudo

	if strings.Contains(c.Request.RequestURI, "?") {
		queryArgs := strings.SplitAfterN(c.Request.RequestURI, "?", 2)[1]
		url += "?" + queryArgs
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	cloneHeaders(c, req)

	resp, err := HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.Status(resp.StatusCode)
		c.Writer.Write(body)
	} else {
		c.JSON(resp.StatusCode, gin.H{"message": "mAPPa Backend error", "status": resp.Status, "error": err})
	}
}

func Ping() (int, string, error) {
	res, err := http.Head(UrlMappa)
	if err == nil {
		return res.StatusCode, res.Status, nil
	}
	return 0, err.Error(), err
}

func LoginStatsRoute(c *gin.Context) {
	stats := GetStats()
	c.JSON(200, stats)
}
func LoginRoute(c *gin.Context) {
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Failed to get request body: %s\n", err)
		c.JSON(400, gin.H{"message": "mAPPa request error", "error": err.Error()})
		return
	}
	var loginRequest LoginRequest
	err = json.Unmarshal(requestBody, &loginRequest)
	if err != nil {
		log.Printf("Login request body is invalid: %s\n", err)
		c.JSON(400, gin.H{"message": "mAPPA request error", "error": err.Error()})
		return
	}
	loginResponse, ok := Login(loginRequest.UserName, loginRequest.Password)
	if !ok {
		c.JSON(403, gin.H{"message": "mAPPa login failed"})
		return
	}
	c.JSON(202, loginResponse)
}

func Login(username string, password string) (loginResponse LoginResponse, ok bool) {
	login, ok := GetLogin(username, password)
	if ok {
		log.Printf("Cached login recovered for user %s", username)
		return login, ok
	}
	loginResponse, ok = PostLogin(username, password)
	if ok {
		SetLogin(username, password, loginResponse)
	}
	return loginResponse, ok
}

func cloneHeaders(c *gin.Context, req *http.Request) {
	allowedHeaders := []string{"Authorization", "User-Agent", "Host"}
	for _, s := range allowedHeaders {
		headerValue := c.GetHeader(s)
		if len(headerValue) > 0 {
			req.Header.Set(s, headerValue)
		}
	}
}
