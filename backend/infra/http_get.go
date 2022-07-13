package infra

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	netUrl "net/url"
)

func HttpGet(url string, headers map[string]string, logMessage string) (statusCode int, responseBody []byte, err error) {
	defer func() {
		if err != nil {			
			log.Printf("GET %s - ERROR %v", url, err)
		}
	}()

	escapedUrl := URLEscape(url)
	if escapedUrl != url {
		log.Printf("HTTP GET %s -> %s", url, escapedUrl)
	}

	req, err := http.NewRequest("GET", escapedUrl, nil)

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

func URLEscape(url string) (escapedUrl string) {
	parsedUrl, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return url // TODO: Implementar controle de erro
	}
	query := ""
	for key, value := range parsedUrl.Query() {
		if len(query) > 0 {
			query += "&"
		}
		query += netUrl.QueryEscape(key) + "=" + netUrl.QueryEscape(value[0])
	}
	if len(query) > 0 {
		escapedUrl = fmt.Sprintf("%s://%s%s?%s", parsedUrl.Scheme, parsedUrl.Hostname(), parsedUrl.Path, query)
	} else {
		escapedUrl = fmt.Sprintf("%s://%s%s", parsedUrl.Scheme, parsedUrl.Hostname(), parsedUrl.Path)
	}
	return escapedUrl
}
