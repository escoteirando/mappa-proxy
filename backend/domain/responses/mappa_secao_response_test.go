package responses

import (
	"testing"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/types"
)

const arg = `
[
	{
		"codigo":1,
		"nome":"Seção 1",
		"codigoTipoSecao":1,
		"codigoGrupo":1,
		"codigoRegiao":"teste",
		"subsecoes":[
			{
				"codigo":1,
				"nome":"Subseção 1",
				"codigoSecao":1,
				"codigoLider":1,
				"codigoViceLider":0,
				"associados":[
					{
						"codigo":1,
						"nome":"Associado 1",
						"codigoFoto":0,
						"codigoEquipe":1,
						"username":1,
						"numeroDigito":1,
						"dataNascimento":"2022-07-13T17:29:48-03:00",
						"dataValidade":"2022-07-13T17:29:48-03:00",
						"nomeAbreviado":"A1",
						"sexo":"M",
						"codigoRamo":1,
						"codigoCategoria":1,
						"codigoSegundaCategoria":1,
						"codigoTerceiraCategoria":0,
						"linhaFormacao":"teste",
						"codigoRamoAdulto":1,
						"dataAcompanhamento":"2022-07-13T17:29:48-03:00"
					}
				]
			}
		]
	}
]`

func TestGetMappaSecaoResponsesFromJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	associado := &MappaAssociadoResponse{
		Codigo:                  1,
		Nome:                    "Associado 1",
		CodigoFoto:              0,
		CodigoEquipe:            1,
		UserName:                1,
		NumeroDigito:            1,
		DataNascimento:          types.Date(time.Now()),
		DataValidade:            types.Date(time.Now()),
		NomeAbreviado:           "A1",
		Sexo:                    "M",
		CodigoRamo:              1,
		CodigoCategoria:         1,
		CodigoSegundaCategoria:  1,
		CodigoTerceiraCategoria: 0,
		LinhaFormacao:           "teste",
		CodigoRamoAdulto:        1,
		DataAcompanhamento:      types.Date(time.Now()),
	}
	subSecao := &MappaSubSecaoResponse{
		Codigo:          1,
		Nome:            "Subseção 1",
		CodigoSecao:     1,
		CodigoLider:     1,
		CodigoViceLider: 0,
		Associados:      []*MappaAssociadoResponse{associado},
	}
	secao := &MappaSecaoResponse{
		Codigo:          1,
		Nome:            "Seção 1",
		CodigoTipoSecao: 1,
		CodigoGrupo:     1,
		CodigoRegiao:    "teste",
		Subsecoes:       []*MappaSubSecaoResponse{subSecao},
	}
	sample := []*MappaSecaoResponse{secao}

	// body, _ := json.Marshal(sample)

	tests := []struct {
		name    string
		args    args
		want    []*MappaSecaoResponse
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{[]byte(arg)},
			want:    sample,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMappaSecaoResponsesFromJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMappaSecaoResponsesFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 1 {
				t.Errorf("GetMappaSecaoResponsesFromJSON() expected len = 1, got %d", len(got))
			}
			if len(got[0].Subsecoes) != 1 {
				t.Errorf("GetMappaSecaoResponsesFromJSON() expected len = 1, got %d", len(got[0].Subsecoes))
			}

		})
	}
}
