package responses

type HealthCheckResponse struct {
	Status      string              `json:"status"`
	MappaServer MappaServerResponse `json:"mappa_server"`
}