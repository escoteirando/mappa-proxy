package infra

import (
	"log"
	"net/http"
	"testing"
)

func TestHttpRequestNetHttp_HttpGet(t *testing.T) {
	type fields struct {
		HttpClient *http.Client
		logger     *log.Logger
	}
	type args struct {
		url        string
		headers    map[string]string
		logMessage string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantStatusCode   int
		wantResponseBody []byte
		wantErr          bool
	}{
		{
			name: "GET request Successfully",
			args: args{
				url:        "https://httpbin.org/get",
				headers:    map[string]string{"Content-Type": "application/json"},
				logMessage: "GET request Successfully",
			},
			wantStatusCode: 200,
		},
		{
			name: "GET request Failed",
			args: args{
				url:        "https://httpbin.org/get_ERROR",
				headers:    map[string]string{"Content-Type": "application/json"},
				logMessage: "GET request Failed",
			},
			wantStatusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetHttpRequester()
			gotStatusCode, _, err := r.HttpGet(tt.args.url, tt.args.headers, tt.args.logMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpRequestNetHttp.HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("HttpRequestNetHttp.HttpGet() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}

		})
	}
}

func TestHttpRequestNetHttp_HttpPost(t *testing.T) {
	type fields struct {
		HttpClient *http.Client
		logger     *log.Logger
	}
	type args struct {
		url        string
		headers    map[string]string
		logMessage string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantStatusCode   int
		wantResponseBody []byte
		wantErr          bool
	}{
		{
			name: "POST request Successfully",
			args: args{
				url:        "https://httpbin.org/post",
				headers:    map[string]string{"Content-Type": "application/json"},
				logMessage: "POST request Successfully",
			},
			wantStatusCode: 200,
		},
		{
			name: "POST request Failed",
			args: args{
				url:        "https://httpbin.org/post_ERROR",
				headers:    map[string]string{"Content-Type": "application/json"},
				logMessage: "POST request Failed",
			},
			wantStatusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetHttpRequester()
			gotStatusCode, _, err := r.HttpPost(tt.args.url, tt.args.headers, []string{"test"}, tt.args.logMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpRequestNetHttp.HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("HttpRequestNetHttp.HttpGet() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}

		})
	}
}
