package mappa

import (
	"fmt"
	"log"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/infra"
)

func (m *MappaAPI) get(url string, authorization string) (statusCode int, responseBody []byte, err error) {
	startTime := time.Now()
	reqMsg := "Mappa GET " + url
	defer func() {
		log.Printf("%s: %d %v [%v - %dB]", reqMsg, statusCode, err, time.Since(startTime), len(responseBody))
	}()
	requester := infra.GetHttpRequester()
	url = fmt.Sprintf("%s%s", m.mappaUrl, url)
	headers := map[string]string{"Authorization": authorization}

	statusCode, responseBody, err = requester.HttpGet(url, headers, reqMsg)
	statusCode, err = validateHttpResponse(reqMsg, statusCode, responseBody, err)

	return
}

func (m *MappaAPI) post(url string, body interface{}) (statusCode int, responseBody []byte, err error) {
	startTime := time.Now()
	reqMsg := "Mappa POST " + url
	defer func() {
		log.Printf("%s: %d %v [%v - %dB]", reqMsg, statusCode, err, time.Since(startTime), len(responseBody))
	}()
	requester := infra.GetHttpRequester()
	url = fmt.Sprintf("%s%s", m.mappaUrl, url)
	headers := make(map[string]string, 0)
	statusCode, responseBody, err = requester.HttpPost(url, headers, body, reqMsg)
	statusCode, err = validateHttpResponse(reqMsg, statusCode, responseBody, err)
	return
}

func validateHttpResponse(reqMsg string, statusCode int, responseBody []byte, err error) (int, error) {
	if statusCode < 100 || statusCode > 299 {
		err = fmt.Errorf("%s: %d %s", reqMsg, statusCode, string(responseBody))
	}
	return statusCode, err
}
