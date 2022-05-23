package responses

import "encoding/json"

type (
	MappaProgressaoResponse struct {
		Codigo                int    `json:"codigo"`
		Descricao             string `json:"descricao"`
		CodigoUEB             string `json:"codigoUeb"`
		Ordenacao             int    `json:"ordenacao"`
		CodigoCaminho         int    `json:"codigoCaminho"`
		CodigoDesenvolvimento int    `json:"codigoDesenvolvimento"`
		NumeroGrupo           int    `json:"numeroGrupo"`
		CodigoRegiao          string `json:"codigoRegiao"`
		CodigoCompetencia     int    `json:"codigoCompetencia"`
		Segmento              string `json:"segmento"`
	}
	MappaProgressoesResponse []*MappaProgressaoResponse
)

// {
// 		"codigo": 1,
// 		"descricao": "Ouvir o episódio \"Irmãos de Mowgli\" do Livro da Selva.",
// 		"codigoUeb": "S2",
// 		"ordenacao": 2,
// 		"codigoCaminho": 1,
// 		"codigoDesenvolvimento": 23,
// 		"numeroGrupo": null,
// 		"codigoRegiao": null,
// 		"codigoCompetencia": 38,
// 		"segmento": "PROMESSA_ESCOTEIRA_LOBINHO"
// 	},

func GetMappaProgressaoResponsesFromJSON(data []byte) ([]*MappaProgressaoResponse, error) {
	var response []*MappaProgressaoResponse
	err := json.Unmarshal(data, &response)
	return response, err
}
