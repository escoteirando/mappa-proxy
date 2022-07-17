package repositories

import (
	"fmt"
	"path"
	"testing"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
)

func TestDBRepository_SetAssociadoSecao(t *testing.T) {
	tmp := path.Join(t.TempDir(), "db_repository_associado_secao_test.db")
	conn := fmt.Sprintf("sqlite://%s", tmp)

	t.Run("Test associado secao e subsecao", func(t *testing.T) {
		repo := RepositoryFactory.GetRepository(conn)
		if repo == nil {
			t.Error("Repository not created")
		}

		// Criar secoes
		for i := 1; i <= 3; i++ {
			repo.SetSecao(&entities.Secao{
				Codigo:          uint(i),
				Nome:            fmt.Sprintf("Secao %d", i),
				CodigoRegiao:    "SC",
				CodigoTipoSecao: 1,
				CodigoGrupo:     32,
			})

			repo.SetSubSecao(&entities.SubSecao{
				Codigo:          uint(i * 10),
				Nome:            fmt.Sprintf("SubSecao %d", i*10),
				CodigoSecao:     uint(i),
				CodigoLider:     1,
				CodigoViceLider: 2,
			})
		}

		err := repo.SetAssociadoSecao(4, 1, 2, 3)
		if err != nil {
			t.Errorf("Failed to set associado secao: %v", err)
		}

		if err = repo.SetAssociadoSubSecao(4, 10); err != nil {
			t.Errorf("Failed to set associado subsecao: %v", err)
		}

		secoes, err := repo.GetAssociadoSecoes(4)
		if err != nil {
			t.Errorf("Failed to get associado secoes: %v", err)
		}
		if len(secoes) != 3 {
			t.Errorf("Expected 3 secoes, got %d", len(secoes))
		}

		subsecao, err := repo.GetAssociadoSubSecao(4)
		if err != nil {
			t.Errorf("Failed to get associado subsecao: %v", err)
		}
		if subsecao.Nome != "SubSecao 10" {
			t.Errorf("Expected SubSecao 10, got %s", subsecao.Nome)
		}

	})

}
