package responses

type MappaDetalhesResponse struct {
	Escotista *MappaEscotistaResponse `json:"escotista,omitempty"`
	Associado *MappaAssociadoResponse `json:"associado,omitempty"`
	Grupos    *MappaGrupoResponse     `json:"grupos,omitempty"`
}
