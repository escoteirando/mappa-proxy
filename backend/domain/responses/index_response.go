package responses

type IndexResponse struct {
	App       string `json:"app"`
	Version   string `json:"version"`
	RunningBy string `json:"running-by"`
}
