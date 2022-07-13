package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpPost(url string, data interface{}, logMessage string) (statusCode int, responseBody []byte, err error) {
	defer func() {
		if err != nil {
			log.Printf("POST %s - ERROR %v", url, err)
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
	resp, err := HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	responseBody, err = ioutil.ReadAll(resp.Body)
	return
}
