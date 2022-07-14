package responses

import "encoding/json"

type MappaSecaoResponse struct {
	Codigo          uint                     `json:"codigo"`
	Nome            string                   `json:"nome"`
	CodigoTipoSecao uint                     `json:"codigoTipoSecao"`
	CodigoGrupo     uint                     `json:"codigoGrupo"`
	CodigoRegiao    string                   `json:"codigoRegiao"`
	Subsecoes       []*MappaSubSecaoResponse `json:"subsecoes"`
}

func GetMappaSecaoResponsesFromJSON(data []byte) ([]*MappaSecaoResponse, error) {
	var response []*MappaSecaoResponse
	err := json.Unmarshal(data, &response)
	return response, err
}
