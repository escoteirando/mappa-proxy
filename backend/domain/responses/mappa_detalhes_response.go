package responses

type MappaDetalhesResponse struct {
	Escotista *MappaEscotistaResponse `json:"escotista,omitempty"`
	Associado *MappaAssociadoResponse `json:"associado,omitempty"`
	Grupo     *MappaGrupoResponse     `json:"grupo,omitempty"`
}
