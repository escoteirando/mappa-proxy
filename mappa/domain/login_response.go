package domain

import "time"

type LoginResponse struct {
	ID      string    `json:"id"`
	TTL     int       `json:"ttl"`
	Created time.Time `json:"created"`
	Userid  int       `json:"userId"`
}
