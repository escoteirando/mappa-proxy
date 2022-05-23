package types

import (
	"reflect"
	"testing"
)

func TestBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		value   bool
		want    []byte
		wantErr bool
	}{
		{
			name:    "true",
			value:   true,
			want:    []byte("true"),
			wantErr: false,
		}, {
			name:    "false",
			value:   false,
			want:    []byte("false"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bool{
				value: tt.value,
			}
			got, err := b.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bool.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value bool
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		json    string
		value   Bool
		wantErr bool
	}{
		{
			name:    "true",
			json:    `"true"`,
			value:   Bool{value: true},
			wantErr: false,
		}, {
			name:    "false",
			json:    `"false"`,
			value:   Bool{value: false},
			wantErr: false,
		},
		{
			name:    "Sim",
			json:    `"S"`,
			value:   Bool{value: true},
			wantErr: false,
		}, {
			name:    "Não",
			json:    `"N"`,
			value:   Bool{value: false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bool{}
			if err := b.UnmarshalJSON([]byte(tt.json)); (err != nil) != tt.wantErr || tt.value.value != b.value {
				t.Errorf("Bool.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
