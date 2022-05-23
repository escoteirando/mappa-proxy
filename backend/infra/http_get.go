package infra

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet(url string, headers map[string]string, logMessage string) (statusCode int, responseBody []byte, err error) {
	defer func() {
		if err == nil {
			log.Print(logMessage)
		} else {
			log.Printf("GET %s - ERROR %v", url, err)
		}
	}()
	req, err := http.NewRequest("GET", url, nil)
	for header, value := range headers {
		req.Header.Add(header, value)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", "okhttp/3.4.1")
	resp, err := HttpClient.Do(req)
	statusCode = resp.StatusCode
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	return
}
