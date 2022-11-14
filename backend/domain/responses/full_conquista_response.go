package responses

import "time"

type (
	FullConquistaResponse struct {
		Type             string                     `json:"type"`
		Data             time.Time                  `json:"dataConquista"`
		CodAssociado     int                        `json:"codigoAssociado"`
		CodEscotista     int                        `json:"codigoEscotistaUltimaAlteracao"`
		NumeroNivel      int                        `json:"numeroNivel"`
		CodEspecialidade int                        `json:"codigoEspecialidade"`
		Associado        *MappaAssociadoResponse     `json:"associado"`
		Especialidade    *MappaEspecialidadeResponse `json:"especialidade"`
	}
)
