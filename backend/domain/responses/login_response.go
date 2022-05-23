package responses

import "time"

type MappaLoginResponse struct {
	ID       string                 `json:"id"`
	TTL      int                    `json:"ttl"`
	Created  time.Time              `json:"created"`
	Userid   int                    `json:"userId"`
}

func (loginResponse *MappaLoginResponse) ValidUntil() time.Time {
	return loginResponse.Created.Add(time.Duration(loginResponse.TTL) * time.Second)
}

func (loginResponse *MappaLoginResponse) IsValid() bool {
	return loginResponse.ValidUntil().After(time.Now())
}
