package domain

// LoginRequest {"type":"LOGIN_REQUEST","username":"guionardo","password":"****"}
type LoginRequest struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
