package responses

type MappaServerResponse struct {
	MappaServerUrl string `json:"mappa_server_url"`
	StatusCode     int    `json:"status_code"`
	Status         string `json:"status"`
}
