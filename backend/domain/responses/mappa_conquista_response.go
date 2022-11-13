package responses

import "github.com/escoteirando/mappa-proxy/backend/types"

type (
	MappaConquistaResponse struct {
		Type             string     `json:"type"`
		Data             types.Time `json:"dataConquista"`
		CodAssociado     int        `json:"codigoAssociado"`
		CodEscotista     int        `json:"codigoEscotistaUltimaAlteracao"`
		NumeroNivel      int        `json:"numeroNivel"`
		CodEspecialidade int        `json:"codigoEspecialidade"`
	}
	MappaConquistasResponse struct {
		DataHora types.Time                `json:"dataHora"`
		Values   []*MappaConquistaResponse `json:"values"`
	}
)

//  {
//         "type": "SERVICOS",
//         "dataConquista": "2022-08-20T13:56:40.000Z",
//         "codigoAssociado": 1107587,
//         "codigoEscotistaUltimaAlteracao": 50442,
//         "numeroNivel": 1,
//         "codigoEspecialidade": 20
//     }
