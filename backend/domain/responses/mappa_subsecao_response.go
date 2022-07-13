package responses

import "encoding/json"

type MappaSubSecaoResponse struct {
	Codigo          uint                      `json:"codigo"`
	Nome            string                    `json:"nome"`
	CodigoSecao     uint                      `json:"codigoSecao"`
	CodigoLider     uint                      `json:"codigoLider"`
	CodigoViceLider uint                      `json:"codigoViceLider"`
	Associados      []*MappaAssociadoResponse `json:"associados"`
}

func GetSubSecoesFromJSON(data []byte) ([]MappaSubSecaoResponse, error) {
	var response []MappaSubSecaoResponse
	err := json.Unmarshal(data, &response)
	return response, err
}