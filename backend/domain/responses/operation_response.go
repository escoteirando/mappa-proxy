package responses

import "github.com/escoteirando/mappa-proxy/backend/infra"

type OperationResponse struct {
	HealthCheck  HealthCheckResponse `json:"health_check"`
	MemoryStatus infra.MemoryStatus  `json:"memory"`
}
