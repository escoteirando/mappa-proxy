package responses

import "github.com/escoteirando/mappa-proxy/backend/types"

type (
	MappaMarcacaoResponse struct {
		CodigoAtividade       int        `json:"codigoAtividade"`
		CodigoAssociado       int        `json:"codigoAssociado"`
		DataAtividade         types.Time `json:"dataAtividade"`
		DataStatusEscotista   types.Time `json:"dataStatusEscotista"`
		DataHoraAtualizacao   types.Time `json:"dataHoraAtualizacao"`
		CodigoUltimoEscotista int        `json:"codigoUltimoEscotista"`
		Segmento              string     `json:"segmento"`
	}
	MappaMarcacoesResponse struct {
		DataHora types.Time               `json:"dataHora"`
		Values   []*MappaMarcacaoResponse `json:"values"`
	}
)
