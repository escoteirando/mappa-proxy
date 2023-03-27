package responses

import "github.com/escoteirando/mappa-proxy/backend/infra"

type HealthCheckResponse struct {
	Status      string              `json:"status"`
	MappaServer MappaServerResponse `json:"mappa_server"`
	Memory      infra.MemoryStatus  `json:"memory"`
}
