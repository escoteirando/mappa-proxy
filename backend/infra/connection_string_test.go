package infra

import (
	"reflect"
	"testing"
)

func TestCreateConnectionString(t *testing.T) {
	type args struct {
		connectionString string
	}
	tests := []struct {
		name    string
		args    args
		want    *ConnectionString
		wantErr bool
	}{
		{
			name:    "valid sqlite",
			args:    args{connectionString: "sqlite:test.db"},
			want:    &ConnectionString{Schema: "sqlite", ConnectionData: "test.db"},
			wantErr: false,
		},
		{
			name:    "invalid",
			args:    args{connectionString: "invalid"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateConnectionString(tt.args.connectionString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConnectionString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
