package responses

import "encoding/json"

type (
	MappaEspecialidadeResponse struct {		
		Codigo           int                              `json:"codigo"`
		Descricao        string                           `json:"descricao"`
		RamoConhecimento string                           `json:"ramoConhecimento"`
		PreRequisito     string                           `json:"prerequisito"`
		Itens            []MappaEspecialidadeItemResponse `json:"itens"`
	}
	MappaEspecialidadeItemResponse struct {
		// Id                  int    `json:"id"`
		CodigoEspecialidade int    `json:"codigoEspecialidade"`
		Descricao           string `json:"descricao"`
		Numero              int    `json:"numero"`
	}
)

func GetMappaEspecialidadesResponseFromJSON(data []byte) ([]*MappaEspecialidadeResponse, error) {
	var response []*MappaEspecialidadeResponse
	err := json.Unmarshal(data, &response)
	return response, err
}
