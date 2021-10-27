package mappa

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guionardo/mappa_proxy/mappa/domain"
	"github.com/guionardo/mappa_proxy/mappa/repositories"
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

func Login(username string, password string) (loginResponse domain.LoginResponse, ok bool) {
	repository := repositories.GetRepository()
	login, err := repository.GetLogin(username, password)
	if err {
		log.Printf("Cached login recovered for user %s", username)
		return login, ok
	}

	loginResponse, ok = PostLogin(username, password)
	if ok {
		repository.SetLogin(username, password, loginResponse, time.Now())
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
