package responses

import (
	"encoding/json"

	"github.com/escoteirando/mappa-proxy/backend/types"
)

type MappaEscotistaResponse struct {
	UserId          int        `json:"codigo"`
	CodigoAssociado int        `json:"codigoAssociado"`
	UserName        string     `json:"username"`
	NomeCompleto    string     `json:"nomeCompleto"`
	Ativo           types.Bool `json:"ativo"`
	CodigoGrupo     int        `json:"codigoGrupo"`
	CodigoRegiao    string     `json:"codigoRegiao"`
	CodigoFoto      int        `json:"codigoFoto"`

	// 	{
	// 	"codigo": 50442,
	// 	"codigoAssociado": 850829,
	// 	"username": "Guionardo",
	// 	"nomeCompleto": "GuionardoFurlan",
	// 	"ativo": "S",
	// 	"codigoGrupo": 32,
	// 	"codigoRegiao": "SC",
	// 	"codigoFoto": null
	// }
}

func GetMappaEscotistaResponseFromJSON(data []byte) *MappaEscotistaResponse {
	var response MappaEscotistaResponse
	json.Unmarshal(data, &response)
	return &response
}
