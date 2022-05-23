package responses

import "encoding/json"

type MappaEspecialidadeResponse struct {
	// TODO: Implementar estrutura de dados Especialidade
}

func GetMappaEspecialidadesResponseFromJSON(data []byte) (*MappaEspecialidadeResponse, error) {
	var response MappaEspecialidadeResponse
	err := json.Unmarshal(data, &response)
	return &response, err
}
