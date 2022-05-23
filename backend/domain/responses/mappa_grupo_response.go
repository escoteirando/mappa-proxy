package responses

import "encoding/json"

type MappaGrupoResponse struct {
	Codigo           int    `json:"codigo"`
	CodigoRegiao     string `json:"codigoRegiao"`
	Nome             string `json:"nome"`
	CodigoModalidade int    `json:"codigoModalidade"`
	// 	[
	// 	{
	// 		"codigo": 32,
	// 		"codigoRegiao": "SC",
	// 		"nome": "LEÕES DE BLUMENAU",
	// 		"codigoModalidade": 1
	// 	}
	// ]
}

func GetMappaGruposResponseFromJSON(data []byte) []*MappaGrupoResponse {
	var response []*MappaGrupoResponse
	json.Unmarshal(data, &response)
	return response
}
