package mappa

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

// GET /api/marcacoes/v2/updates?dataHoraUltimaAtualizacao=2019-08-31T13:48:23-00:00&codigoSecao=1424 HTTP/1.1
func (api *MappaAPI) GetMarcacoes(codigoSecao int, desde time.Time, authorization string) (marcacoes *responses.MappaMarcacoesResponse, err error) {

	_, body, err := api.get(fmt.Sprintf("/api/marcacoes/v2/updates?dataHoraUltimaAtualizacao=%s&codigoSecao=%d", desde.Format(time.RFC3339), codigoSecao), authorization)

	if err == nil {
		err = json.Unmarshal(body, &marcacoes)
	}
	return

}
