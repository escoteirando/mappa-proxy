package responses

type IndexResponse struct {
	App       string `json:"app"`
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`	
	RunningBy string `json:"running_by"`
}
