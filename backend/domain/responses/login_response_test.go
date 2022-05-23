package responses

import (
	"encoding/json"
	"testing"
)

func TestLoginResponse_IsValid(t *testing.T) {
	loginResponseRaw := `{
		"id":"yVkd6s8wyTWFxQjVlO0TYUAdBUwbq7G0pd4OJBWjK8CAcXgxI9YdVHcw0aZqcV6t",
		"ttl":1209600,"created":"2021-08-06T17:47:20.642Z",
		"userId":50442}`

	t.Run("default", func(t *testing.T) {
		var loginResponse *MappaLoginResponse
		if err := json.Unmarshal([]byte(loginResponseRaw), &loginResponse); err != nil {
			t.Errorf("Failed to unmarshal %v", err)
		}
		validUntil := loginResponse.ValidUntil()
		if loginResponse.IsValid() {
			t.Errorf("Expected login response is invalid - %v", validUntil)
		}
	})

}
