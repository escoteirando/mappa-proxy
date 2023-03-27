package infra

import "testing"

func TestURLEscape(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		wantEscapedUrl string
	}{
		{
			name:           "URL without query",
			url:            "https://httpbin.org/get",
			wantEscapedUrl: "https://httpbin.org/get",
		},
		{
			name:           "URL with query",
			url:            "https://httpbin.org/get?param1=value1&param2=value2",
			wantEscapedUrl: "https://httpbin.org/get?param1=value1&param2=value2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEscapedUrl := URLEscape(tt.url); gotEscapedUrl != tt.wantEscapedUrl {
				t.Errorf("URLEscape() = %v, want %v", gotEscapedUrl, tt.wantEscapedUrl)
			}
		})
	}
}
