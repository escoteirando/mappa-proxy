package infra

import (
	"log"
	"os"
	"strings"
)

type (
	HttpRequestMock struct {
		setups map[string]RequestMock
		logger *log.Logger
	}
	RequestMock struct {
		Method             string
		Url                string
		Data               interface{}
		ResponseStatusCode int
		ResponseBody       []byte
		Err                error
	}
)

func NewHttpRequestMock() *HttpRequestMock {
	return &HttpRequestMock{
		logger: log.New(os.Stdout, "[HttpRequestMock] ", log.LstdFlags),
		setups: make(map[string]RequestMock),
	}
}
func (r *HttpRequestMock) Setup(mock RequestMock) *HttpRequestMock {
	r.setups[strings.ToUpper(mock.Method)+":"+URLEscape(mock.Url)] = mock
	return r
}

func (r *HttpRequestMock) HttpGet(url string, headers map[string]string, logMessage string) (statusCode int, responseBody []byte, err error) {
	defer func() {
		if err != nil {
			r.logger.Printf("GET %s - ERROR %v", url, err)
		}
	}()

	escapedUrl := strings.Split(URLEscape(url), "?")[0]

	if response, ok := r.setups["GET:"+escapedUrl]; ok {
		return response.ResponseStatusCode, response.ResponseBody, response.Err
	}
	return 404, []byte("Not Found"), nil
}

func (r *HttpRequestMock) HttpPost(url string, headers map[string]string, data interface{}, logMessage string) (statusCode int, responseBody []byte, err error) {
	defer func() {
		if err != nil {
			r.logger.Printf("POST %s - ERROR %v", url, err)
		}
	}()
	escapedUrl := URLEscape(url)

	if response, ok := r.setups["POST:"+escapedUrl]; ok {
		return response.ResponseStatusCode, response.ResponseBody, response.Err
	}
	return 404, []byte("Not Found"), nil
}
