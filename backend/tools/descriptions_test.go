package tools

import (
	"reflect"
	"testing"
)

func TestGetWordsFromExpression(t *testing.T) {

	tests := []struct {
		expression string
		wantWords  []string
	}{
		{"PROMESSA_ESCOTEIRA_LOBINHO", []string{"Promessa", "Escoteira", "Lobinho"}},
		{"PROGRESSAO_SENIOR", []string{"Progressão", "Sênior"}},
		{"PROGRESSAO_RASTREADOR_CACADOR", []string{"Progressão", "Rastreador", "Caçador"}},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			if gotWords := GetWordsFromExpression(tt.expression); !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("GetWordsFromExpression() = %v, want %v", gotWords, tt.wantWords)
			}
		})
	}
}
