package requests

import (
	"reflect"
	"testing"
)

func TestCreateMappaLoginRequest(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *LoginRequest
	}{
		{
			name: "default",
			args: args{
				username: "guionardo",
				password: "test"},
			want: &LoginRequest{
				Type:     "LOGIN_REQUEST",
				UserName: "guionardo",
				Password: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMappaLoginRequest(tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMappaLoginRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
