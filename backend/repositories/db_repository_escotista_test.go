package repositories

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/types"
)

func TestDBRepository_DetalheEscotista(t *testing.T) {
	tmp := path.Join(t.TempDir(), "db_repository_detalhe_escotista_test.db")
	conn := fmt.Sprintf("sqlite://%s", tmp)
	t.Run("test", func(t *testing.T) {

		repository := RepositoryFactory.GetRepository(conn)
		if repository == nil {
			t.Errorf("Failed to create repository")
		}

		escotista := responses.MappaEscotistaResponse{
			UserId:          1234,
			CodigoAssociado: 2222,
			UserName:        "test_user",
			NomeCompleto:    "TEST USER",
			Ativo:           types.TrueBool,
			CodigoGrupo:     32,
			CodigoRegiao:    "SC",
			CodigoFoto:      0,
		}
		associado := responses.MappaAssociadoResponse{
			Codigo:                  2222,
			Nome:                    "TEST",
			CodigoFoto:              0,
			CodigoEquipe:            12,
			UserName:                1234,
			NumeroDigito:            5,
			DataNascimento:          types.Date(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
			DataValidade:            types.Date(time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC)),
			NomeAbreviado:           "TESTE",
			Sexo:                    "M",
			CodigoRamo:              1,
			CodigoCategoria:         0,
			CodigoSegundaCategoria:  0,
			CodigoTerceiraCategoria: 0,
			LinhaFormacao:           "TESTE",
			CodigoRamoAdulto:        0,
			DataAcompanhamento:      types.Date(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
		}
		grupo := responses.MappaGrupoResponse{
			Codigo:           32,
			CodigoRegiao:     "SC",
			Nome:             "GELB",
			CodigoModalidade: 1,
		}
		if err := repository.SetDetalheEscotista(7777, responses.MappaDetalhesResponse{
			Escotista: &escotista,
			Associado: &associado,
			Grupo:     &grupo,
		}); err != nil {
			t.Errorf("DBRepository.SetDetalheEscotista() error = %v", err)
		}

		detalhes, err := repository.GetDetalheEscotista(7777)
		if err != nil {
			t.Errorf("DBRepository.GetDetalheEscotista() error = %v", err)
		}
		if detalhes.Escotista.UserName != "test_user" {
			t.Errorf("DBRepository.GetDetalheEscotista() error = %v", err)
		}

	})

}
