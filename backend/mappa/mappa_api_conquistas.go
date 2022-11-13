package mappa

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

// /api/associado-conquistas/v2/updates?dataHoraUltimaAtualizacao=2020-09-05T15:04:59-00:00&codigoSecao=1424&codigoEscotista=850829
func (api *MappaAPI) GetConquistas(codigoSecao int, desde time.Time, authorization string) (conquistas *responses.MappaConquistasResponse, err error) {
	_, body, err := api.get(fmt.Sprintf("/api/associado-conquistas/v2/updates?dataHoraUltimaAtualizacao=%s&codigoSecao=%d", desde.Format(time.RFC3339), codigoSecao), authorization)

	if err == nil {
		err = json.Unmarshal(body, &conquistas)
	}

	return
}
