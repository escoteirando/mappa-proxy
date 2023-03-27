package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HttpRequestNetHttp struct {
	HttpClient *http.Client
	logger     *log.Logger
}

func NewHttpRequest() *HttpRequestNetHttp {
	requester := &HttpRequestNetHttp{
		HttpClient: &http.Client{},
		logger:     log.New(os.Stdout, "[HttpRequest] ", log.LstdFlags),
	}

	requester.logger.Printf("Initialized")
	return requester
}

func (r *HttpRequestNetHttp) HttpGet(url string, headers map[string]string, logMessage string) (statusCode int, responseBody []byte, err error) {
	defer func() {
		if err != nil {
			r.logger.Printf("GET %s - ERROR %v", url, err)
		}
	}()

	escapedUrl := URLEscape(url)
	if escapedUrl != url {
		r.logger.Printf("HTTP GET %s -> %s", url, escapedUrl)
	}

	req, err := http.NewRequest("GET", escapedUrl, nil)

	for header, value := range headers {
		req.Header.Add(header, value)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", "okhttp/3.4.1")
	resp, err := r.HttpClient.Do(req)
	statusCode = resp.StatusCode
	if err != nil {
		r.logger.Fatal(err)
	}
	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	return
}

func (r *HttpRequestNetHttp) HttpPost(url string, headers map[string]string, data interface{}, logMessage string) (statusCode int, responseBody []byte, err error) {
	defer func() {
		if err != nil {
			r.logger.Printf("POST %s - ERROR %v", url, err)
		}
	}()
	body, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("Failed to serialize %v - %v", data, err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		err = fmt.Errorf("Failed to create request %s", err)
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", "okhttp/3.4.1")
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		r.logger.Printf("POST EXCEPTION %v", err)
		return
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	responseBody, err = ioutil.ReadAll(resp.Body)
	return
}
