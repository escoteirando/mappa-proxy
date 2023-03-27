package responses

type StatsResponse struct {
	NumeroEscotistas int `json:"escotistas"`
	NumeroAssociados int `json:"associados"`
	NumeroGrupos     int `json:"grupos"`
	NumeroSecoes     int `json:"secoes"`
}
