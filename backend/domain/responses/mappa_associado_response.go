package responses

import (
	"encoding/json"

	"github.com/escoteirando/mappa-proxy/backend/types"
)

type MappaAssociadoResponse struct {
	Codigo                  int        `json:"codigo"`
	Nome                    string     `json:"nome"`
	CodigoFoto              int        `json:"codigoFoto"`
	CodigoEquipe            int        `json:"codigoEquipe"`
	UserName                int        `json:"username"`
	NumeroDigito            int        `json:"numeroDigito"`
	DataNascimento          types.Time `json:"dataNascimento"`
	DataValidade            types.Time `json:"dataValidade"`
	NomeAbreviado           string     `json:"nomeAbreviado"`
	Sexo                    string     `json:"sexo"`
	CodigoRamo              int        `json:"codigoRamo"`
	CodigoCategoria         int        `json:"codigoCategoria"`
	CodigoSegundaCategoria  int        `json:"codigoSegundaCategoria"`
	CodigoTerceiraCategoria int        `json:"codigoTerceiraCategoria"`
	LinhaFormacao           string     `json:"linhaFormacao"`
	CodigoRamoAdulto        int        `json:"codigoRamoAdulto"`
	DataAcompanhamento      types.Time `json:"dataAcompanhamento"`

	// 	"codigo": 850829,
	// 	"nome": "GUIONARDO FURLAN",
	// 	"codigoFoto": null,
	// 	"codigoEquipe": null,
	// 	"username": 1247937,
	// 	"numeroDigito": 3,
	// 	"dataNascimento": "Sat Feb 05 1977 00:00:00 GMT+0000 (Coordinated Universal Time)",
	// 	"dataValidade": "2022-01-01T00:00:00.000Z",
	// 	"nomeAbreviado": "GUIONARDO",
	// 	"sexo": "M",
	// 	"codigoRamo": 2,
	// 	"codigoCategoria": 5,
	// 	"codigoSegundaCategoria": 0,
	// 	"codigoTerceiraCategoria": 0,
	// 	"linhaFormacao": "Escotista",
	// 	"codigoRamoAdulto": 2,
	// 	"dataAcompanhamento": null
	// }
}

func GetMappaAssociadoResponseFromJSON(data []byte) *MappaAssociadoResponse {
	var response MappaAssociadoResponse
	json.Unmarshal(data, &response)
	return &response
}
