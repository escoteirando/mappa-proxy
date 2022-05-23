package requests

// LoginRequest {"type":"LOGIN_REQUEST","username":"guionardo","password":"****"}
type LoginRequest struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func CreateMappaLoginRequest(username string, password string) *LoginRequest {
	return &LoginRequest{
		Type:     "LOGIN_REQUEST",
		UserName: username,
		Password: password,
	}
}
