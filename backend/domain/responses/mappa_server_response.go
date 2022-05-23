package responses

import (
	"github.com/escoteirando/mappa-proxy/backend/infra"
)

type MappaServerResponse struct {
	MappaServerUrl string             `json:"mappa_server_url"`
	StatusCode     int                `json:"status_code"`
	Status         string             `json:"status"`
	Memory         infra.MemoryStatus `json:"memory"`
}
