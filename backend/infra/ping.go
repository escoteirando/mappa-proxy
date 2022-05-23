package infra

import "net/http"

func Ping(url string) (int, string, error) {
	res, err := http.Head(url)
	if err == nil {
		return res.StatusCode, res.Status, nil
	}
	return 0, err.Error(), err
}
