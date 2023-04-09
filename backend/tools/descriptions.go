package tools

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var translations map[string]string

// Recebe expressões como PROMESSA_ESCOTEIRA_LOBINHO e
// retorna ["Promessa","Escoteira","Lobinho"]
func GetWordsFromExpression(expression string) (words []string) {
	words = strings.Split(expression, "_")
	for i := 0; i < len(words); i++ {
		if w, ok := translations[words[i]]; ok {
			words[i] = w
		} else {
			words[i] = cases.Title(language.BrazilianPortuguese).String(words[i])
		}

	}
	return
}

func GetPhraseFromExpression(expression string) string {
	words := GetWordsFromExpression(expression)
	return strings.Join(words, " ")
}

func init() {
	translations = make(map[string]string, 10)
	translations["CAMPEOES"] = "Campeões"
	translations["CACADOR"] = "Caçador"
	translations["ACAO"] = "Ação"
	translations["SERVICOS"] = "Serviços"
	translations["CIENCIA"] = "Ciência"
	translations["PROGRESSAO"] = "Progressão"
	translations["PATATENRA"] = "Pata-tenra"
	translations["SENIOR"] = "Sênior"
	translations["CARATER"] = "Caráter"
	translations["FISICO"] = "Físico"
}
