package responses

import (
	"reflect"
	"testing"

	"github.com/escoteirando/mappa-proxy/backend/types"
)

func TestGetMappaEscotistaResponseFromJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want *MappaEscotistaResponse
	}{
		{
			name: "Valid",
			args: args{
				data: []byte(`{"codigo":50442,"codigoAssociado":850829,"username":"Guionardo","nomeCompleto":"GuionardoFurlan","ativo":"S","codigoGrupo":32,"codigoRegiao":"SC","codigoFoto":null}`),
			},
			want: &MappaEscotistaResponse{
				UserId:          50442,
				CodigoAssociado: 850829,
				UserName:        "Guionardo",
				NomeCompleto:    "GuionardoFurlan",
				Ativo:           types.TrueBool,
				CodigoGrupo:     32,
				CodigoRegiao:    "SC",
				// CodigoFoto:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMappaEscotistaResponseFromJSON(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMappaEscotistaResponseFromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
