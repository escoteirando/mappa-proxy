package infra

import (
	"fmt"
	netUrl "net/url"
)

type HttpRequester interface {
	HttpGet(url string, headers map[string]string, logMessage string) (statusCode int, responseBody []byte, err error)
	HttpPost(url string, headers map[string]string, data interface{}, logMessage string) (statusCode int, responseBody []byte, err error)
}

var httpRequester HttpRequester

func GetHttpRequester() HttpRequester {
	if httpRequester == nil {
		httpRequester = NewHttpRequest()
	}
	return httpRequester
}

func ResetHttpRequester(requester HttpRequester) {
	httpRequester = requester
}

func URLEscape(url string) (escapedUrl string) {
	parsedUrl, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return url // TODO: Implementar controle de erro
	}
	query := ""
	if queries := parsedUrl.Query(); len(queries) > 0 {
		query = "?" + queries.Encode()
		// for key,value:=range queries{
		// query += netUrl.QueryEscape(key) + "=" + netUrl.QueryEscape(value[0])
		// }
	}

	// for key, value := range parsedUrl.Query() {
	// 	if len(query) > 0 {
	// 		query += "&"
	// 	}
	// 	query += netUrl.QueryEscape(key) + "=" + netUrl.QueryEscape(value[0])
	// }
	escapedUrl = fmt.Sprintf("%s://%s%s%s", parsedUrl.Scheme, parsedUrl.Host, parsedUrl.Path, query)
	// if len(query) > 0 {
	// 	escapedUrl = fmt.Sprintf("%s://%s%s?%s", parsedUrl.Scheme, parsedUrl.Hostname(), parsedUrl.Path, query)
	// } else {
	// 	escapedUrl = fmt.Sprintf("%s://%s%s", parsedUrl.Scheme, parsedUrl.Hostname(), parsedUrl.Path)
	// }
	return escapedUrl
}
