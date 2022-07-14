package repositories

import (
	"fmt"
	"path"
	"testing"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
)

func TestDBRepository_SetGrupo(t *testing.T) {
	tmp := path.Join(t.TempDir(), "db_repository_grupo_test.db")
	conn := fmt.Sprintf("sqlite://%s", tmp)

	t.Run("Test Grupo", func(t *testing.T) {
		repo := RepositoryFactory.GetRepository(conn)
		if repo == nil {
			t.Error("Repository not created")
		}
		grupo := &entities.Grupo{
			Codigo:           32,
			CodigoRegiao:     "SC",
			Nome:             "GELB",
			CodigoModalidade: 1,
		}
		if err := repo.SetGrupo(grupo); err != nil {
			t.Errorf("Error setting grupo: %v", err)
		}

		grupo2, err := repo.GetGrupo(32)
		if err != nil {
			t.Errorf("Error getting grupo: %v", err)
		}

		if grupo.Codigo != grupo2.Codigo || grupo.CodigoRegiao != grupo2.CodigoRegiao || grupo.Nome != grupo2.Nome || grupo.CodigoModalidade != grupo2.CodigoModalidade {
			t.Errorf("Grupos not equal: %v != %v", grupo, grupo2)
		}

	})

}
